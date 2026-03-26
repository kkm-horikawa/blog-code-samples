package main

import (
	"fmt"
	"sync"
	"time"
)

// === 基本のゴルーチン ===
func sayHello(name string, ch chan<- string) {
	time.Sleep(100 * time.Millisecond) // 何か処理
	ch <- fmt.Sprintf("こんにちは、%sさん!", name)
}

// === ワーカーパターン ===
func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		time.Sleep(50 * time.Millisecond) // 処理を模擬
		results <- fmt.Sprintf("ワーカー%d がジョブ%d を完了", id, job)
	}
}

func main() {
	// === 基本: ゴルーチン + チャネル ===
	fmt.Println("--- ゴルーチンとチャネル ---")
	ch := make(chan string)

	go sayHello("太郎", ch)
	go sayHello("花子", ch)

	// 2つの結果を受信
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// === バッファ付きチャネル ===
	fmt.Println("\n--- バッファ付きチャネル ---")
	buffered := make(chan int, 3)
	buffered <- 10
	buffered <- 20
	buffered <- 30
	fmt.Println(<-buffered) // 10（FIFO）
	fmt.Println(<-buffered) // 20
	fmt.Println(<-buffered) // 30

	// === WaitGroup で全ゴルーチンの完了を待つ ===
	fmt.Println("\n--- WaitGroup ---")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			time.Sleep(time.Duration(n*50) * time.Millisecond)
			fmt.Printf("ゴルーチン %d 完了\n", n)
		}(i)
	}

	wg.Wait() // 全部終わるまで待つ
	fmt.Println("全ゴルーチン完了")

	// === select で複数チャネルを待つ ===
	fmt.Println("\n--- select ---")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "ch1からのメッセージ"
	}()
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "ch2からのメッセージ"
	}()

	// 先に届いた方を2回受信
	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}

	// === ワーカープール ===
	fmt.Println("\n--- ワーカープール ---")
	jobs := make(chan int, 10)
	results := make(chan string, 10)

	var workerWg sync.WaitGroup
	// 3つのワーカーを起動
	for w := 1; w <= 3; w++ {
		workerWg.Add(1)
		go worker(w, jobs, results, &workerWg)
	}

	// 5つのジョブを投入
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // これ以上ジョブがないことを通知

	// ワーカー完了後に results を閉じる
	go func() {
		workerWg.Wait()
		close(results)
	}()

	// 結果を受信
	for result := range results {
		fmt.Println(result)
	}
}
