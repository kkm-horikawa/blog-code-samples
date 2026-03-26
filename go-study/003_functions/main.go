package main

import "fmt"

// 基本の関数
func add(a, b int) int {
	return a + b
}

// 複数戻り値
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("0で割れません")
	}
	return a / b, nil
}

// 名前付き戻り値
func rect(w, h float64) (area, perimeter float64) {
	area = w * h
	perimeter = 2 * (w + h)
	return
}

// 可変長引数
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// 関数を引数に取る（高階関数）
func apply(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func main() {
	// 基本の関数呼び出し
	fmt.Printf("3 + 5 = %d\n", add(3, 5))

	// 複数戻り値の受け取り
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", result)
	}

	// エラーになるケース
	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("エラー:", err)
	}

	// 名前付き戻り値
	area, perimeter := rect(5, 3)
	fmt.Printf("面積: %.1f, 周長: %.1f\n", area, perimeter)

	// 可変長引数
	fmt.Printf("合計: %d\n", sum(1, 2, 3, 4, 5))

	// 高階関数
	multiply := func(a, b int) int { return a * b }
	fmt.Printf("4 * 6 = %d\n", apply(4, 6, multiply))
}
