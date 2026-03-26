package main

import (
	"fmt"
	"sort"
)

func main() {
	// === スライスの作成 ===
	nums := []int{10, 20, 30, 40, 50}
	fmt.Println("元のスライス:", nums)

	// === append で追加 ===
	nums = append(nums, 60)
	fmt.Println("追加後:", nums)

	// === スライシング（部分取得） ===
	fmt.Println("nums[1:3]:", nums[1:3]) // [20, 30]
	fmt.Println("nums[:3]:", nums[:3])   // [10, 20, 30]
	fmt.Println("nums[3:]:", nums[3:])   // [40, 50, 60]

	// === len と cap ===
	s := make([]int, 3, 10)
	fmt.Printf("len=%d, cap=%d, 値=%v\n", len(s), cap(s), s)

	// === range で回す ===
	fruits := []string{"バナナ", "りんご", "みかん", "ぶどう"}
	for i, f := range fruits {
		fmt.Printf("[%d] %s\n", i, f)
	}

	// === ソート ===
	sort.Strings(fruits)
	fmt.Println("ソート後:", fruits)

	numbers := []int{5, 3, 8, 1, 9, 2}
	sort.Ints(numbers)
	fmt.Println("数値ソート:", numbers)

	// === フィルタリングのパターン ===
	// 偶数だけ抽出
	all := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	even := []int{} // 空スライスから始める
	for _, n := range all {
		if n%2 == 0 {
			even = append(even, n)
		}
	}
	fmt.Println("偶数:", even)

	// === スライスは参照型 ===
	original := []int{1, 2, 3}
	copied := original          // コピーではなく同じ配列を参照
	copied[0] = 999             // original も変わる!
	fmt.Println("original:", original) // [999, 2, 3]

	// 本当にコピーしたい場合
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	dst[0] = 999
	fmt.Println("src:", src) // [1, 2, 3] 変わらない
	fmt.Println("dst:", dst) // [999, 2, 3]
}
