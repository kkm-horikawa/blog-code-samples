package main

import "fmt"

// パッケージレベルの変数（var のみ使える）
var greeting = "こんにちは"

// 定数
const pi = 3.14159

func main() {
	// var で明示的に型を指定
	var name string = "太郎"
	var age int = 25
	var isStudent bool = true

	// := で短縮宣言（型推論）
	city := "東京"
	score := 98.5

	fmt.Println(greeting)
	fmt.Printf("名前: %s\n", name)
	fmt.Printf("年齢: %d\n", age)
	fmt.Printf("学生: %t\n", isStudent)
	fmt.Printf("都市: %s\n", city)
	fmt.Printf("点数: %.1f\n", score)
	fmt.Printf("円周率: %f\n", pi)

	// ゼロ値の確認
	var x int
	var s string
	var b bool
	fmt.Printf("\nゼロ値: int=%d, string=%q, bool=%t\n", x, s, b)

	// 複数変数の同時宣言
	var (
		width  = 100
		height = 200
	)
	fmt.Printf("面積: %d\n", width*height)
}
