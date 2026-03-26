package main

import (
	"errors"
	"fmt"
)

// === センチネルエラー（パッケージレベルで定義） ===
var (
	ErrNotFound   = errors.New("見つかりません")
	ErrOutOfRange = errors.New("範囲外です")
)

// === カスタムエラー型 ===
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("バリデーションエラー [%s]: %s", e.Field, e.Message)
}

// === エラーを返す関数 ===
func findUser(id int) (string, error) {
	users := map[int]string{
		1: "太郎",
		2: "花子",
	}

	name, ok := users[id]
	if !ok {
		// エラーをラップして文脈を追加
		return "", fmt.Errorf("ユーザーID=%d: %w", id, ErrNotFound)
	}
	return name, nil
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{
			Field:   "age",
			Message: fmt.Sprintf("年齢は0以上である必要があります（入力値: %d）", age),
		}
	}
	if age > 150 {
		return &ValidationError{
			Field:   "age",
			Message: fmt.Sprintf("年齢は150以下である必要があります（入力値: %d）", age),
		}
	}
	return nil
}

// === defer のデモ ===
func countDown() {
	fmt.Println("\nカウントダウン開始")
	for i := 3; i >= 1; i-- {
		defer fmt.Printf("  defer: %d\n", i) // LIFO順で実行
	}
	fmt.Println("カウントダウン終了")
	// ここで defer が 1, 2, 3 の逆順で実行される
}

// === panic と recover ===
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("パニックから回復: %v", r)
		}
	}()

	return a / b, nil // b=0 だとパニック
}

func main() {
	// === 基本のエラーハンドリング ===
	name, err := findUser(1)
	if err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Println("見つかりました:", name)
	}

	// エラーケース
	_, err = findUser(99)
	if err != nil {
		fmt.Println("エラー:", err)

		// errors.Is でセンチネルエラーを判定
		if errors.Is(err, ErrNotFound) {
			fmt.Println("→ NotFoundエラーです")
		}
	}

	// === カスタムエラー型の判定 ===
	fmt.Println("\n--- バリデーション ---")
	if err := validateAge(-5); err != nil {
		fmt.Println("エラー:", err)

		// errors.As でカスタムエラー型を取得
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Printf("→ フィールド: %s\n", ve.Field)
		}
	}

	if err := validateAge(25); err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Println("年齢25: OK")
	}

	// === defer ===
	countDown()

	// === panic と recover ===
	fmt.Println("\n--- panic/recover ---")
	result, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Println("結果:", result)
	}
}
