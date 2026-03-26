package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// === データモデル ===
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// === インメモリストア（簡易DB代わり） ===
var (
	users  = []User{{ID: 1, Name: "太郎", Email: "taro@example.com"}}
	nextID = 2
	mu     sync.Mutex
)

// === ハンドラ関数 ===

// ユーザー一覧
func listUsers(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, users)
}

// ユーザー取得（パスパラメータ）
func getUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id") // Go 1.22+

	for _, u := range users {
		if fmt.Sprintf("%d", u.ID) == id {
			writeJSON(w, http.StatusOK, u)
			return
		}
	}
	writeJSON(w, http.StatusNotFound, map[string]string{
		"error": "ユーザーが見つかりません",
	})
}

// ユーザー作成
func createUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "不正なリクエストボディ",
		})
		return
	}

	if input.Name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "name は必須です",
		})
		return
	}

	mu.Lock()
	user := User{ID: nextID, Name: input.Name, Email: input.Email}
	users = append(users, user)
	nextID++
	mu.Unlock()

	writeJSON(w, http.StatusCreated, user)
}

// === ミドルウェア ===

// ログミドルウェア
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// CORS ミドルウェア
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// === ヘルパー ===

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// === メイン ===

func main() {
	mux := http.NewServeMux()

	// Go 1.22: メソッド + パスパラメータが標準で使える
	mux.HandleFunc("GET /api/users", listUsers)
	mux.HandleFunc("POST /api/users", createUser)
	mux.HandleFunc("GET /api/users/{id}", getUser)

	// ヘルスチェック
	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	// ミドルウェアを適用
	handler := loggingMiddleware(corsMiddleware(mux))

	fmt.Println("サーバー起動: http://localhost:8080")
	fmt.Println("  GET  /api/health")
	fmt.Println("  GET  /api/users")
	fmt.Println("  POST /api/users")
	fmt.Println("  GET  /api/users/{id}")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
