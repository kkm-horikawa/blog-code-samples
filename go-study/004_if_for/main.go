package main

import "fmt"

func main() {
	// === if / else ===
	age := 25
	if age >= 20 {
		fmt.Println("成人です")
	} else if age >= 13 {
		fmt.Println("ティーンエイジャーです")
	} else {
		fmt.Println("子供です")
	}

	// === if の初期化文 ===
	// 変数のスコープを if ブロック内に限定できる
	if score := 85; score >= 80 {
		fmt.Printf("スコア %d: 合格!\n", score)
	}

	// === for: 通常のカウントループ ===
	fmt.Print("カウント: ")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// === for: while のように使う ===
	count := 3
	for count > 0 {
		fmt.Printf("カウントダウン: %d\n", count)
		count--
	}

	// === for: range でスライスを回す ===
	fruits := []string{"りんご", "みかん", "ぶどう"}
	for i, fruit := range fruits {
		fmt.Printf("[%d] %s\n", i, fruit)
	}

	// インデックス不要なら _ で捨てる
	for _, fruit := range fruits {
		fmt.Println("果物:", fruit)
	}

	// === switch ===
	day := "水"
	switch day {
	case "月", "火", "水", "木", "金":
		fmt.Printf("%s曜日は平日\n", day)
	case "土", "日":
		fmt.Printf("%s曜日は休日\n", day)
	default:
		fmt.Println("不明な曜日")
	}

	// === switch: 条件式なし（if-else チェーンの代替） ===
	temp := 32
	switch {
	case temp >= 35:
		fmt.Println("猛暑日")
	case temp >= 30:
		fmt.Println("真夏日")
	case temp >= 25:
		fmt.Println("夏日")
	default:
		fmt.Println("普通の日")
	}

	// === break と continue ===
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue // 3をスキップ
		}
		if i == 7 {
			break // 7で終了
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println() // 出力: 0 1 2 4 5 6
}
