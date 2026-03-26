package main

import (
	"fmt"
	"math"
)

// === インターフェースの定義 ===
type Shape interface {
	Area() float64
	Perimeter() float64
}

// === Circle: Shape を実装 ===
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// === Rectangle: Shape を実装 ===
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// === インターフェースを受け取る関数 ===
func printShapeInfo(s Shape) {
	fmt.Printf("  面積: %.2f, 周長: %.2f\n", s.Area(), s.Perimeter())
}

// === 小さなインターフェース（Go流） ===
type Speaker interface {
	Speak() string
}

type Dog struct{ Name string }
type Cat struct{ Name string }

func (d Dog) Speak() string { return d.Name + ": ワン!" }
func (c Cat) Speak() string { return c.Name + ": ニャー!" }

// === 空インターフェース / 型アサーション ===
func describe(x any) string {
	switch v := x.(type) {
	case int:
		return fmt.Sprintf("整数: %d", v)
	case string:
		return fmt.Sprintf("文字列: %q", v)
	case bool:
		return fmt.Sprintf("真偽値: %t", v)
	default:
		return fmt.Sprintf("その他: %v (%T)", v, v)
	}
}

func main() {
	// === ポリモーフィズム ===
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 10},
	}

	for _, s := range shapes {
		fmt.Printf("%T:\n", s)
		printShapeInfo(s)
	}

	// === 小さなインターフェース ===
	fmt.Println("\n--- 動物 ---")
	animals := []Speaker{
		Dog{Name: "ポチ"},
		Cat{Name: "タマ"},
		Dog{Name: "ハチ"},
	}
	for _, a := range animals {
		fmt.Println(a.Speak())
	}

	// === 型アサーション ===
	fmt.Println("\n--- 型アサーション ---")
	fmt.Println(describe(42))
	fmt.Println(describe("hello"))
	fmt.Println(describe(true))
	fmt.Println(describe(3.14))

	// 個別の型アサーション
	var x any = "Go言語"
	if str, ok := x.(string); ok {
		fmt.Printf("文字列を取得: %s (長さ: %d)\n", str, len(str))
	}
}
