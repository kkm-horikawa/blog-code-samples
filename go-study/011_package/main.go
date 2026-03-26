package main

import (
	"fmt"

	"example.com/study/greeting" // 自作パッケージのインポート
)

func main() {
	// 公開関数を呼ぶ
	fmt.Println(greeting.Hello("太郎"))
	fmt.Println(greeting.Formal("田中"))

	// greeting.format("太郎") → コンパイルエラー!（非公開）
}
