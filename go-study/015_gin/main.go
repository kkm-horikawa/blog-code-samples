package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

// === データモデル ===
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// リクエスト用（バリデーション付き）
type CreateUserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

// === インメモリストア ===
var (
	users  = []User{{ID: 1, Name: "太郎", Email: "taro@example.com"}}
	nextID = 2
	mu     sync.Mutex
)

// === ハンドラ ===

func listUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"count": len(users),
	})
}

func getUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "不正なID"})
		return
	}

	for _, u := range users {
		if u.ID == id {
			c.JSON(http.StatusOK, u)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "ユーザーが見つかりません"})
}

func createUser(c *gin.Context) {
	var input CreateUserInput

	// ShouldBindJSON: JSONパース + バリデーション
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	user := User{ID: nextID, Name: input.Name, Email: input.Email}
	users = append(users, user)
	nextID++
	mu.Unlock()

	c.JSON(http.StatusCreated, user)
}

func deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "不正なID"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "削除しました"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "ユーザーが見つかりません"})
}

// === ミドルウェア ===

// 簡易認証ミドルウェア
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "Bearer secret-token" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "認証が必要です",
			})
			return
		}
		c.Next()
	}
}

// === メイン ===

func main() {
	r := gin.Default() // Logger + Recovery ミドルウェア付き

	// 公開API
	api := r.Group("/api/v1")
	{
		api.GET("/users", listUsers)
		api.GET("/users/:id", getUser)
		api.POST("/users", createUser)
	}

	// 認証が必要なAPI
	admin := api.Group("", authMiddleware())
	{
		admin.DELETE("/users/:id", deleteUser)
	}

	// ヘルスチェック
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	fmt.Println("gin サーバー起動: http://localhost:8080")
	fmt.Println("  GET    /api/v1/users")
	fmt.Println("  POST   /api/v1/users")
	fmt.Println("  GET    /api/v1/users/:id")
	fmt.Println("  DELETE /api/v1/users/:id  (要認証)")
	r.Run(":8080")
}
