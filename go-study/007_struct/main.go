package main

import "fmt"

// === 構造体の定義 ===
type User struct {
	Name  string
	Age   int
	Email string
}

// === 値レシーバのメソッド（読み取り専用） ===
func (u User) Greet() string {
	return fmt.Sprintf("こんにちは、%sです（%d歳）", u.Name, u.Age)
}

func (u User) String() string {
	return fmt.Sprintf("User{Name: %s, Age: %d, Email: %s}", u.Name, u.Age, u.Email)
}

// === ポインタレシーバのメソッド（変更可能） ===
func (u *User) Birthday() {
	u.Age++
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

// === コンストラクタパターン ===
func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

// === 構造体の埋め込み（コンポジション） ===
type Admin struct {
	User        // User を埋め込み
	Role string
}

func (a Admin) ShowRole() string {
	return fmt.Sprintf("%s は %s です", a.Name, a.Role)
}

func main() {
	// 構造体の作成方法いろいろ
	u1 := User{Name: "太郎", Age: 25, Email: "taro@example.com"}
	u2 := User{"花子", 30, "hanako@example.com"} // フィールド順に指定（非推奨）
	u3 := NewUser("次郎", 20)                      // コンストラクタ

	fmt.Println(u1.Greet())
	fmt.Println(u2)
	fmt.Println(u3)

	// ポインタレシーバでフィールドを変更
	fmt.Printf("\n誕生日前: %d歳\n", u1.Age)
	u1.Birthday()
	fmt.Printf("誕生日後: %d歳\n", u1.Age)

	u3.SetEmail("jiro@example.com")
	fmt.Println("メール設定後:", u3)

	// 構造体の埋め込み
	admin := Admin{
		User: User{Name: "管理者", Age: 35},
		Role: "スーパー管理者",
	}
	// User のメソッドがそのまま使える
	fmt.Println("\n" + admin.Greet())
	fmt.Println(admin.ShowRole())

	// 構造体の比較
	a := User{Name: "太郎", Age: 25}
	b := User{Name: "太郎", Age: 25}
	fmt.Printf("\na == b: %t\n", a == b) // true
}
