package main

import (
	"fmt"
	"sort"
)

func main() {
	// === マップの作成 ===
	scores := map[string]int{
		"国語": 80,
		"数学": 95,
		"英語": 88,
	}
	fmt.Println("成績:", scores)

	// === 要素の追加と更新 ===
	scores["理科"] = 72
	scores["国語"] = 85 // 上書き
	fmt.Println("更新後:", scores)

	// === 要素の取得 ===
	math := scores["数学"]
	fmt.Println("数学:", math)

	// === 存在チェック（カンマOKイディオム） ===
	value, ok := scores["社会"]
	if ok {
		fmt.Println("社会:", value)
	} else {
		fmt.Println("社会の成績はありません")
	}

	// 短縮形
	if v, ok := scores["数学"]; ok {
		fmt.Printf("数学の点数は %d です\n", v)
	}

	// === 要素の削除 ===
	delete(scores, "英語")
	fmt.Println("英語を削除:", scores)

	// === マップの走査（順序は不定） ===
	fmt.Println("\n--- 全科目 ---")
	for subject, score := range scores {
		fmt.Printf("%s: %d点\n", subject, score)
	}

	// === ソートして出力したい場合 ===
	keys := make([]string, 0, len(scores))
	for k := range scores {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Println("\n--- ソート済み ---")
	for _, k := range keys {
		fmt.Printf("%s: %d点\n", k, scores[k])
	}

	// === 文字の出現回数をカウント ===
	text := "hello world"
	charCount := make(map[rune]int)
	for _, ch := range text {
		charCount[ch]++
	}
	fmt.Println("\n文字カウント:")
	for ch, count := range charCount {
		fmt.Printf("  '%c': %d回\n", ch, count)
	}

	// === nil マップに注意 ===
	// var m map[string]int  // nil マップ
	// m["a"] = 1            // パニック!
	// 必ず make か リテラルで初期化する
	m := make(map[string]int)
	m["a"] = 1 // OK
	fmt.Println("\n安全なマップ:", m)
}
