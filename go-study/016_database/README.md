# 016: データベース接続 (GORM)

## このレッスンで学ぶこと

- GORM の基本操作（CRUD）
- モデル定義と構造体タグ
- マイグレーション（AutoMigrate）
- リレーション（1対多）
- トランザクション
- クエリビルダ

## 前提

```bash
go mod init study/016_database
go get gorm.io/gorm
go get gorm.io/driver/sqlite  # 学習用はSQLite
```

本番では `gorm.io/driver/postgres` + `pgx` を使う。

## コード解説

### モデル定義

```go
type User struct {
    gorm.Model          // ID, CreatedAt, UpdatedAt, DeletedAt が自動追加
    Name  string `gorm:"not null"`
    Email string `gorm:"uniqueIndex;not null"`
    Posts []Post  // 1対多リレーション
}
```

`gorm.Model` は Django の `models.Model` に相当。

### CRUD

```go
// Create
db.Create(&User{Name: "太郎", Email: "taro@example.com"})

// Read
var user User
db.First(&user, 1)                          // ID=1
db.Where("name = ?", "太郎").First(&user)    // 条件検索

// Update
db.Model(&user).Update("Name", "次郎")

// Delete（論理削除: DeletedAt に日時が入る）
db.Delete(&user, 1)
```

### マイグレーション

```go
db.AutoMigrate(&User{}, &Post{})
```

Django の `makemigrations` + `migrate` に相当。
ただし、AutoMigrate は**カラム追加のみ**。カラム削除・型変更はしない。
本番では golang-migrate や atlas を使う。

### リレーション

```go
// 1対多: User has many Posts
type Post struct {
    gorm.Model
    Title  string
    Body   string
    UserID uint   // 外部キー（命名規則で自動認識）
    User   User   // belongs to
}

// Preload でリレーションを読み込み
db.Preload("Posts").Find(&users)
```

Django の `ForeignKey` + `select_related` / `prefetch_related` に相当。

## Django ORM との対比

| Django ORM | GORM |
|-----------|------|
| `User.objects.create(name="太郎")` | `db.Create(&User{Name: "太郎"})` |
| `User.objects.get(id=1)` | `db.First(&user, 1)` |
| `User.objects.filter(name="太郎")` | `db.Where("name = ?", "太郎").Find(&users)` |
| `user.save()` | `db.Save(&user)` |
| `user.delete()` | `db.Delete(&user)` |
| `ForeignKey` | 構造体に `UserID uint` + `User User` |
| `select_related` | `db.Preload("Posts")` |
| `makemigrations` / `migrate` | `db.AutoMigrate()` (開発用) |

## やってみよう

1. `Tag` モデルを追加して、Post と多対多リレーションを作ってみよう
2. `db.Where` でいろんな条件検索を試してみよう
3. トランザクション内で複数のレコードを更新してみよう
