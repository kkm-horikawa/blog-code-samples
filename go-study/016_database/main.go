package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// === モデル定義 ===

type User struct {
	gorm.Model        // ID, CreatedAt, UpdatedAt, DeletedAt
	Name  string `gorm:"not null"`
	Email string `gorm:"uniqueIndex;not null"`
	Posts []Post // 1対多リレーション
}

type Post struct {
	gorm.Model
	Title  string `gorm:"not null"`
	Body   string
	UserID uint // 外部キー（User.ID を参照）
	User   User // belongs to
}

func main() {
	// === DB接続（SQLite） ===
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal("DB接続に失敗:", err)
	}
	fmt.Println("DB接続成功")

	// === マイグレーション ===
	db.AutoMigrate(&User{}, &Post{})
	fmt.Println("マイグレーション完了")

	// === Create ===
	fmt.Println("\n--- Create ---")
	user1 := User{Name: "太郎", Email: "taro@example.com"}
	user2 := User{Name: "花子", Email: "hanako@example.com"}
	db.Create(&user1)
	db.Create(&user2)
	fmt.Printf("作成: %s (ID=%d)\n", user1.Name, user1.ID)
	fmt.Printf("作成: %s (ID=%d)\n", user2.Name, user2.ID)

	// 記事を作成（リレーション）
	posts := []Post{
		{Title: "Goを始めました", Body: "Hello World!", UserID: user1.ID},
		{Title: "ポインタが難しい", Body: "& と * の違いとは", UserID: user1.ID},
		{Title: "Djangoとの違い", Body: "ORMが全然違う", UserID: user2.ID},
	}
	db.Create(&posts)
	fmt.Printf("記事を %d 件作成\n", len(posts))

	// === Read ===
	fmt.Println("\n--- Read ---")

	// IDで取得
	var foundUser User
	db.First(&foundUser, 1)
	fmt.Printf("ID=1: %s (%s)\n", foundUser.Name, foundUser.Email)

	// 条件で検索
	var taro User
	db.Where("name = ?", "太郎").First(&taro)
	fmt.Printf("名前検索: %s\n", taro.Name)

	// 全件取得
	var allUsers []User
	db.Find(&allUsers)
	fmt.Printf("全ユーザー: %d 人\n", len(allUsers))

	// === Preload（リレーション読み込み） ===
	fmt.Println("\n--- Preload ---")
	var usersWithPosts []User
	db.Preload("Posts").Find(&usersWithPosts)
	for _, u := range usersWithPosts {
		fmt.Printf("%s の記事 (%d件):\n", u.Name, len(u.Posts))
		for _, p := range u.Posts {
			fmt.Printf("  - %s\n", p.Title)
		}
	}

	// === Update ===
	fmt.Println("\n--- Update ---")
	db.Model(&taro).Update("Name", "太郎(更新済み)")
	fmt.Printf("更新後: %s\n", taro.Name)

	// 複数フィールド更新
	db.Model(&taro).Updates(User{Name: "太郎", Email: "taro-new@example.com"})

	// === 条件付きクエリ ===
	fmt.Println("\n--- 条件付きクエリ ---")
	var taroPosts []Post
	db.Where("user_id = ?", taro.ID).Find(&taroPosts)
	fmt.Printf("太郎の記事: %d 件\n", len(taroPosts))

	// Like検索
	var searchResults []Post
	db.Where("title LIKE ?", "%Go%").Find(&searchResults)
	fmt.Printf("タイトルに「Go」を含む記事: %d 件\n", len(searchResults))

	// === トランザクション ===
	fmt.Println("\n--- トランザクション ---")
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&User{Name: "次郎", Email: "jiro@example.com"}).Error; err != nil {
			return err // ロールバック
		}
		if err := tx.Create(&Post{Title: "トランザクションテスト", UserID: 3}).Error; err != nil {
			return err // ロールバック
		}
		return nil // コミット
	})
	if err != nil {
		fmt.Println("トランザクション失敗:", err)
	} else {
		fmt.Println("トランザクション成功")
	}

	// === Delete（論理削除） ===
	fmt.Println("\n--- Delete ---")
	db.Delete(&Post{}, 1) // ID=1 の記事を論理削除
	var remainingPosts []Post
	db.Find(&remainingPosts)
	fmt.Printf("残り記事数: %d\n", len(remainingPosts))

	// 論理削除されたものも含めて取得
	var allPosts []Post
	db.Unscoped().Find(&allPosts)
	fmt.Printf("論理削除含む全記事数: %d\n", len(allPosts))
}
