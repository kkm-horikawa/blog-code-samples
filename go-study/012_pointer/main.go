package main

import "fmt"

// === 値渡し: 元は変わらない ===
func double(n int) {
	n = n * 2
	fmt.Printf("  関数内: %d\n", n)
}

// === ポインタ渡し: 元が変わる ===
func doublePtr(n *int) {
	*n = *n * 2
	fmt.Printf("  関数内: %d\n", *n)
}

// === swap: ポインタの典型例 ===
func swap(a, b *int) {
	*a, *b = *b, *a
}

// === 構造体とポインタ ===
type Config struct {
	Host    string
	Port    int
	Debug   bool
}

// 値レシーバ: コピーに対して動作（元は変わらない）
func (c Config) Summary() string {
	return fmt.Sprintf("%s:%d (debug=%t)", c.Host, c.Port, c.Debug)
}

// ポインタレシーバ: 元の構造体を変更できる
func (c *Config) EnableDebug() {
	c.Debug = true
}

// === new でポインタを作る ===
func newConfig(host string, port int) *Config {
	return &Config{
		Host: host,
		Port: port,
	}
}

// === nil ポインタの安全な扱い ===
func printName(name *string) {
	if name == nil {
		fmt.Println("名前が設定されていません")
		return
	}
	fmt.Println("名前:", *name)
}

func main() {
	// === ポインタの基本 ===
	fmt.Println("--- ポインタの基本 ---")
	x := 42
	p := &x // p は x のアドレス

	fmt.Printf("x の値: %d\n", x)
	fmt.Printf("x のアドレス: %p\n", &x)
	fmt.Printf("p の値（アドレス）: %p\n", p)
	fmt.Printf("*p（デリファレンス）: %d\n", *p)

	*p = 100
	fmt.Printf("*p = 100 後の x: %d\n", x) // x も 100 に変わる

	// === 値渡し vs ポインタ渡し ===
	fmt.Println("\n--- 値渡し vs ポインタ渡し ---")
	n := 10

	fmt.Printf("呼び出し前: %d\n", n)
	double(n)
	fmt.Printf("値渡し後: %d（変わらない）\n", n)

	doublePtr(&n)
	fmt.Printf("ポインタ渡し後: %d（変わった!）\n", n)

	// === swap ===
	fmt.Println("\n--- swap ---")
	a, b := 1, 2
	fmt.Printf("swap前: a=%d, b=%d\n", a, b)
	swap(&a, &b)
	fmt.Printf("swap後: a=%d, b=%d\n", a, b)

	// === 構造体とポインタ ===
	fmt.Println("\n--- 構造体とポインタ ---")
	cfg := newConfig("localhost", 8080)
	fmt.Println("初期:", cfg.Summary())

	cfg.EnableDebug() // ポインタレシーバ → 元が変わる
	fmt.Println("Debug有効化後:", cfg.Summary())

	// === nil ポインタ ===
	fmt.Println("\n--- nil ポインタ ---")
	var namePtr *string
	printName(namePtr) // nil

	name := "太郎"
	printName(&name) // 値あり

	// === スライス・マップは内部的に参照型 ===
	fmt.Println("\n--- スライスは参照型 ---")
	nums := []int{1, 2, 3}
	modify(nums)
	fmt.Println("関数呼び出し後:", nums) // [999, 2, 3] 変わっている
}

func modify(s []int) {
	s[0] = 999 // ポインタ不要で元が変わる
}
