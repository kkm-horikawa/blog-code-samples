package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// === 基本の構造体 ===
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age,omitempty"`  // ゼロ値なら省略
	Password string `json:"-"`             // JSON に含めない
}

// === ネストした構造体 ===
type Article struct {
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Author    User     `json:"author"`
	Tags      []string `json:"tags"`
	CreatedAt JSONTime `json:"created_at"`
}

// === カスタム型: 日時フォーマット ===
type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := time.Time(t).Format("2006-01-02 15:04:05")
	return json.Marshal(formatted)
}

func (t *JSONTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parsed, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}
	*t = JSONTime(parsed)
	return nil
}

// === API レスポンス汎用型 ===
type APIResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func main() {
	// === Marshal: 構造体 → JSON ===
	fmt.Println("--- Marshal ---")
	user := User{
		ID:       1,
		Name:     "太郎",
		Email:    "taro@example.com",
		Age:      0, // omitempty で省略される
		Password: "secret123",
	}

	data, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(data))
	// Password は json:"-" なので出力されない
	// Age は omitempty で省略される

	// === Unmarshal: JSON → 構造体 ===
	fmt.Println("\n--- Unmarshal ---")
	jsonStr := `{
		"id": 2,
		"name": "花子",
		"email": "hanako@example.com",
		"age": 25,
		"unknown_field": "これは無視される"
	}`

	var parsed User
	if err := json.Unmarshal([]byte(jsonStr), &parsed); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("パース結果: %+v\n", parsed)
	// unknown_field は構造体にないので無視される

	// === ネストした構造体 ===
	fmt.Println("\n--- ネスト ---")
	article := Article{
		Title:     "Goを始めました",
		Body:      "Hello World から始めて...",
		Author:    User{ID: 1, Name: "太郎", Email: "taro@example.com"},
		Tags:      []string{"Go", "入門", "プログラミング"},
		CreatedAt: JSONTime(time.Now()),
	}

	data, _ = json.MarshalIndent(article, "", "  ")
	fmt.Println(string(data))

	// === ネストしたJSONのパース ===
	fmt.Println("\n--- ネストJSONのパース ---")
	nestedJSON := `{
		"title": "ポインタの解説",
		"body": "ポインタとは...",
		"author": {"id": 1, "name": "太郎", "email": "taro@example.com"},
		"tags": ["Go", "ポインタ"],
		"created_at": "2025-01-15 10:30:00"
	}`

	var parsedArticle Article
	if err := json.Unmarshal([]byte(nestedJSON), &parsedArticle); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("タイトル: %s\n", parsedArticle.Title)
	fmt.Printf("著者: %s\n", parsedArticle.Author.Name)
	fmt.Printf("タグ: %v\n", parsedArticle.Tags)

	// === 動的なJSON（構造が不明な場合） ===
	fmt.Println("\n--- 動的JSON ---")
	dynamicJSON := `{"name": "太郎", "scores": [90, 85, 78], "active": true}`

	var result map[string]any
	json.Unmarshal([]byte(dynamicJSON), &result)

	fmt.Printf("型: name=%T, scores=%T, active=%T\n",
		result["name"], result["scores"], result["active"])

	// 型アサーションで取り出す
	if name, ok := result["name"].(string); ok {
		fmt.Println("名前:", name)
	}

	// === API レスポンスパターン ===
	fmt.Println("\n--- API レスポンス ---")

	// 成功
	success := APIResponse{
		Status: "ok",
		Data:   []User{{ID: 1, Name: "太郎"}, {ID: 2, Name: "花子"}},
	}
	data, _ = json.MarshalIndent(success, "", "  ")
	fmt.Println(string(data))

	// エラー
	errResp := APIResponse{
		Status: "error",
		Error:  "ユーザーが見つかりません",
	}
	data, _ = json.MarshalIndent(errResp, "", "  ")
	fmt.Println(string(data))
}
