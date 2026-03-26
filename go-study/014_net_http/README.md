# 014: 標準 net/http で Web API

## このレッスンで学ぶこと

- `http.HandleFunc` でルーティング
- リクエストの読み取り（パス、クエリ、ボディ）
- JSON レスポンスの返し方
- ミドルウェアパターン
- Go 1.22 のメソッドベースルーティング

## コード解説

### 最小の HTTP サーバー

```go
http.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello!")
})
http.ListenAndServe(":8080", nil)
```

これだけで Web サーバーが動く。Django の urls.py + views.py に相当。

### JSON レスポンス

```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message": "hello",
    })
}
```

DRF の `Response(data)` に相当する処理を手動で書く。

### Go 1.22 のルーティング強化

```go
mux.HandleFunc("GET /users", listUsers)
mux.HandleFunc("POST /users", createUser)
mux.HandleFunc("GET /users/{id}", getUser)  // パスパラメータ
```

Go 1.22 から**メソッド指定**と**パスパラメータ `{id}`** が標準で使える。

### ミドルウェア

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}
```

Django のミドルウェアと同じ概念。リクエストの前後に処理を挟む。

## 実行方法

```bash
go run main.go
# 別ターミナルで
curl http://localhost:8080/api/users
curl -X POST -d '{"name":"太郎","email":"taro@example.com"}' http://localhost:8080/api/users
curl http://localhost:8080/api/users/1
```

## Django との対比

| Django | Go net/http |
|--------|------------|
| `urls.py` の `path()` | `mux.HandleFunc("GET /path", handler)` |
| `views.py` の関数 | `func(w http.ResponseWriter, r *http.Request)` |
| `JsonResponse(data)` | `json.NewEncoder(w).Encode(data)` |
| `request.GET["key"]` | `r.URL.Query().Get("key")` |
| `request.body` → serializer | `json.NewDecoder(r.Body).Decode(&data)` |
| ミドルウェアクラス | 関数でラップ `func(next http.Handler) http.Handler` |

## やってみよう

1. `/api/health` エンドポイントを追加してみよう
2. クエリパラメータ `?name=太郎` を受け取って返すAPIを作ってみよう
3. 認証ミドルウェア（固定トークンチェック）を追加してみよう
