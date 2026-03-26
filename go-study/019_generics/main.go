package main

import (
	"cmp"
	"fmt"
	"strings"
)

// === 基本のジェネリック関数 ===

// Max: 順序比較可能な型の最大値
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min: 順序比較可能な型の最小値
func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// === スライス操作 ===

// Contains: スライスに要素が含まれるか
func Contains[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

// Filter: 条件に合う要素だけ抽出
func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := []T{}
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map: 各要素を変換
func Map[T any, U any](slice []T, transform func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = transform(v)
	}
	return result
}

// Reduce: スライスを集約
func Reduce[T any, U any](slice []T, initial U, reducer func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = reducer(result, v)
	}
	return result
}

// === カスタム型制約 ===

type Number interface {
	int | int64 | float64
}

func Sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

// === ジェネリックな構造体 ===

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

// === ジェネリックなマップ操作 ===

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func main() {
	// === Max / Min ===
	fmt.Println("--- Max / Min ---")
	fmt.Println("Max(3, 7):", Max(3, 7))
	fmt.Println("Max(3.14, 2.71):", Max(3.14, 2.71))
	fmt.Println("Max(\"apple\", \"banana\"):", Max("apple", "banana"))
	fmt.Println("Min(3, 7):", Min(3, 7))

	// === Contains ===
	fmt.Println("\n--- Contains ---")
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Contains(nums, 3):", Contains(nums, 3))
	fmt.Println("Contains(nums, 9):", Contains(nums, 9))

	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("Contains(fruits, \"banana\"):", Contains(fruits, "banana"))

	// === Filter ===
	fmt.Println("\n--- Filter ---")
	even := Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println("偶数:", even)

	longFruits := Filter(fruits, func(s string) bool { return len(s) > 5 })
	fmt.Println("6文字以上:", longFruits)

	// === Map ===
	fmt.Println("\n--- Map ---")
	doubled := Map(nums, func(n int) int { return n * 2 })
	fmt.Println("2倍:", doubled)

	upper := Map(fruits, strings.ToUpper)
	fmt.Println("大文字:", upper)

	// 型変換: int → string
	strs := Map(nums, func(n int) string { return fmt.Sprintf("#%d", n) })
	fmt.Println("文字列化:", strs)

	// === Reduce ===
	fmt.Println("\n--- Reduce ---")
	total := Reduce(nums, 0, func(acc, n int) int { return acc + n })
	fmt.Println("合計:", total)

	// === Sum（カスタム制約） ===
	fmt.Println("\n--- Sum ---")
	fmt.Println("Sum(int):", Sum([]int{1, 2, 3, 4, 5}))
	fmt.Println("Sum(float64):", Sum([]float64{1.1, 2.2, 3.3}))

	// === Stack（ジェネリック構造体） ===
	fmt.Println("\n--- Stack ---")
	intStack := NewStack[int]()
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	fmt.Printf("スタックサイズ: %d\n", intStack.Len())

	for intStack.Len() > 0 {
		val, _ := intStack.Pop()
		fmt.Printf("  Pop: %d\n", val)
	}

	strStack := NewStack[string]()
	strStack.Push("Go")
	strStack.Push("は")
	strStack.Push("楽しい")
	for strStack.Len() > 0 {
		val, _ := strStack.Pop()
		fmt.Printf("  Pop: %s\n", val)
	}

	// === マップ操作 ===
	fmt.Println("\n--- Keys / Values ---")
	scores := map[string]int{"国語": 80, "数学": 95, "英語": 88}
	fmt.Println("Keys:", Keys(scores))
	fmt.Println("Values:", Values(scores))
}
