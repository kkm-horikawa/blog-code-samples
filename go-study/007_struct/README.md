# 007: 構造体とメソッド

## このレッスンで学ぶこと

- 構造体（`struct`）の定義
- フィールドへのアクセス
- メソッドの定義（値レシーバ / ポインタレシーバ）
- コンストラクタパターン
- 構造体の埋め込み（継承の代替）

## コード解説

### 構造体の定義

```go
type User struct {
    Name  string
    Age   int
    Email string
}
```

Go にクラスはない。**構造体**がデータをまとめる唯一の方法。

### メソッド

```go
// 値レシーバ: 構造体のコピーに対して動作
func (u User) Greet() string {
    return "Hi, " + u.Name
}

// ポインタレシーバ: 元の構造体を変更できる
func (u *User) SetAge(age int) {
    u.Age = age
}
```

- 読み取りだけなら**値レシーバ** `(u User)`
- フィールドを変更するなら**ポインタレシーバ** `(u *User)`

### コンストラクタパターン

Go にコンストラクタ構文はないが、`New` で始まる関数で代替する。

```go
func NewUser(name string, age int) *User {
    return &User{Name: name, Age: age}
}
```

### 構造体の埋め込み

```go
type Admin struct {
    User          // User を埋め込み
    Permissions []string
}
```

継承ではなく**コンポジション**。`Admin` は `User` のフィールドとメソッドをそのまま使える。

## やってみよう

1. `Book` 構造体（タイトル、著者、ページ数）を作り、情報を出力するメソッドを付けよう
2. ポインタレシーバでページ数を更新するメソッドを追加しよう
3. `Library` 構造体に `[]Book` を持たせて、本の追加・検索機能を作ってみよう
