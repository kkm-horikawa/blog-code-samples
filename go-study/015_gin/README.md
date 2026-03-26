# 015: gin で Web API

## このレッスンで学ぶこと

- gin のセットアップとルーティング
- パスパラメータ・クエリパラメータ
- リクエストボディのバインド（バリデーション付き）
- ミドルウェアとグループルーティング
- エラーハンドリング

## 前提

```bash
go mod init study/015_gin
go get github.com/gin-gonic/gin
```

## コード解説

### 基本構造

```go
r := gin.Default() // Logger + Recovery ミドルウェア付き

r.GET("/users", listUsers)
r.POST("/users", createUser)
r.GET("/users/:id", getUser)

r.Run(":8080")
```

`gin.Default()` でロガーとパニックリカバリが自動適用される。
Django の `DEBUG=True` + ミドルウェアに相当。

### リクエストボディのバインド

```go
type CreateUserInput struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
}

func createUser(c *gin.Context) {
    var input CreateUserInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    // input.Name, input.Email が使える
}
```

構造体タグの `binding` でバリデーション。DRF の Serializer に相当。

### gin.Context

gin の中心は `*gin.Context`。リクエスト情報の取得とレスポンスの返却を一手に担う。

```go
c.Param("id")            // パスパラメータ /users/:id
c.Query("page")           // クエリパラメータ ?page=1
c.DefaultQuery("page", "1")
c.ShouldBindJSON(&input)  // JSONボディをパース
c.JSON(200, data)          // JSONレスポンス
c.Set("userID", 123)       // コンテキストに値を保存
c.Get("userID")            // コンテキストから値を取得
```

### グループルーティング

```go
api := r.Group("/api/v1")
{
    api.GET("/users", listUsers)
    api.POST("/users", createUser)
}

admin := r.Group("/admin", authMiddleware())
{
    admin.DELETE("/users/:id", deleteUser)
}
```

Django の `include()` に相当。ミドルウェアもグループ単位で適用できる。

## net/http との比較

| net/http | gin |
|----------|-----|
| `http.HandleFunc` | `r.GET` / `r.POST` |
| `r.PathValue("id")` | `c.Param("id")` |
| `json.NewDecoder(r.Body).Decode(&v)` | `c.ShouldBindJSON(&v)` |
| `json.NewEncoder(w).Encode(v)` | `c.JSON(200, v)` |
| 手動でミドルウェアをラップ | `r.Use(middleware)` |

## やってみよう

1. `/api/v1/users/:id` の PUT（更新）と DELETE（削除）を追加してみよう
2. ページネーション（`?page=1&limit=10`）を実装してみよう
3. 認証ミドルウェアで `Authorization` ヘッダをチェックしてみよう
