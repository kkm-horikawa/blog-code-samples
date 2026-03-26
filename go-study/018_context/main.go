package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// === タイムアウト付きの処理 ===
func slowQuery(ctx context.Context, query string) (string, error) {
	// ランダムに 1〜5 秒かかるクエリを模擬
	duration := time.Duration(1+rand.Intn(5)) * time.Second

	select {
	case <-time.After(duration):
		// 処理完了
		return fmt.Sprintf("結果: %s (%v)", query, duration), nil
	case <-ctx.Done():
		// タイムアウト or キャンセル
		return "", fmt.Errorf("クエリ中断: %w", ctx.Err())
	}
}

// === context を受け取る関数（Go の慣習） ===
func fetchData(ctx context.Context, id int) (string, error) {
	// ctx のキャンセルを確認してから処理開始
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	// 実際の処理
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("data-%d", id), nil
}

// === 複数ゴルーチンの一斉キャンセル ===
func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("  ワーカー%d: キャンセルされました (%v)\n", id, ctx.Err())
			return
		case <-time.After(200 * time.Millisecond):
			fmt.Printf("  ワーカー%d: 処理中...\n", id)
		}
	}
}

func main() {
	// === 1. context.WithTimeout ===
	fmt.Println("=== WithTimeout（3秒制限） ===")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // 必ず呼ぶ

	result, err := slowQuery(ctx, "SELECT * FROM users")
	if err != nil {
		fmt.Println("タイムアウト:", err)
	} else {
		fmt.Println("成功:", result)
	}

	// === 2. context.WithCancel ===
	fmt.Println("\n=== WithCancel（手動キャンセル） ===")

	ctx2, cancel2 := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx2, i, &wg)
	}

	// 800ms 後にキャンセル
	time.Sleep(800 * time.Millisecond)
	fmt.Println("キャンセル発行!")
	cancel2()

	wg.Wait()
	fmt.Println("全ワーカー停止")

	// === 3. context.WithDeadline ===
	fmt.Println("\n=== WithDeadline（期限指定） ===")

	deadline := time.Now().Add(2 * time.Second)
	ctx3, cancel3 := context.WithDeadline(context.Background(), deadline)
	defer cancel3()

	fmt.Printf("期限: %v\n", deadline.Format("15:04:05"))

	result, err = slowQuery(ctx3, "SELECT * FROM posts")
	if err != nil {
		fmt.Println("期限切れ:", err)
	} else {
		fmt.Println("成功:", result)
	}

	// === 4. context.WithValue（値の伝搬） ===
	fmt.Println("\n=== WithValue（値の伝搬） ===")

	type contextKey string
	const requestIDKey contextKey = "requestID"

	ctx4 := context.WithValue(context.Background(), requestIDKey, "req-abc-123")

	// 深い関数呼び出しの先で取得
	processRequest(ctx4, requestIDKey)

	// === 5. 並列処理 + タイムアウト ===
	fmt.Println("\n=== 並列フェッチ + タイムアウト ===")

	ctx5, cancel5 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel5()

	results := make(chan string, 3)
	for i := 1; i <= 3; i++ {
		go func(id int) {
			data, err := fetchData(ctx5, id)
			if err != nil {
				results <- fmt.Sprintf("ID=%d: エラー %v", id, err)
				return
			}
			results <- fmt.Sprintf("ID=%d: %s", id, data)
		}(i)
	}

	for i := 0; i < 3; i++ {
		fmt.Println(" ", <-results)
	}
}

func processRequest(ctx context.Context, key any) {
	if reqID, ok := ctx.Value(key).(string); ok {
		fmt.Printf("リクエストID: %s\n", reqID)
	}
}
