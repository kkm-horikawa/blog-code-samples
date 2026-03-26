package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// === 1. os.WriteFile / os.ReadFile（簡易版） ===
	fmt.Println("--- WriteFile / ReadFile ---")

	content := "こんにちは、Go!\nファイル操作の学習です。\n3行目のテキスト。\n"
	err := os.WriteFile("sample.txt", []byte(content), 0644)
	if err != nil {
		fmt.Println("書き込みエラー:", err)
		return
	}
	fmt.Println("sample.txt を作成しました")

	data, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("読み込みエラー:", err)
		return
	}
	fmt.Printf("内容:\n%s\n", string(data))

	// === 2. bufio.Scanner（行ごとに読む） ===
	fmt.Println("--- bufio.Scanner ---")

	f, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("  %d: %s\n", lineNum, scanner.Text())
		lineNum++
	}

	// === 3. io.Reader の統一性 ===
	fmt.Println("\n--- io.Reader の統一性 ---")

	// 文字列 → Reader
	sr := strings.NewReader("文字列からの読み込み")
	printFromReader("strings.Reader", sr)

	// バッファ → Reader
	buf := bytes.NewBufferString("バッファからの読み込み")
	printFromReader("bytes.Buffer", buf)

	// ファイル → Reader
	f2, _ := os.Open("sample.txt")
	defer f2.Close()
	printFromReader("os.File", f2)

	// === 4. io.Copy（Reader → Writer） ===
	fmt.Println("\n--- io.Copy ---")

	src := strings.NewReader("io.Copy でコピーされたテキスト")
	dst := &bytes.Buffer{}
	written, _ := io.Copy(dst, src)
	fmt.Printf("コピーしたバイト数: %d\n", written)
	fmt.Printf("コピー結果: %s\n", dst.String())

	// === 5. io.MultiWriter（複数に同時書き込み） ===
	fmt.Println("\n--- io.MultiWriter ---")

	var buf1, buf2 bytes.Buffer
	multi := io.MultiWriter(&buf1, &buf2)
	fmt.Fprintln(multi, "これは2つの Writer に同時に書かれる")

	fmt.Printf("buf1: %s", buf1.String())
	fmt.Printf("buf2: %s", buf2.String())

	// === 6. io.TeeReader（読みながらコピー） ===
	fmt.Println("\n--- io.TeeReader ---")

	original := strings.NewReader("TeeReader のデモ")
	var copied bytes.Buffer
	tee := io.TeeReader(original, &copied)

	// tee から読むと、同時に copied にも書かれる
	result, _ := io.ReadAll(tee)
	fmt.Printf("読み取り: %s\n", string(result))
	fmt.Printf("コピー:   %s\n", copied.String())

	// === 7. bufio.Writer（バッファ付き書き込み） ===
	fmt.Println("\n--- bufio.Writer ---")

	outFile, _ := os.Create("output.txt")
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	for i := 1; i <= 5; i++ {
		fmt.Fprintf(writer, "行 %d: Go のバッファ付き書き込み\n", i)
	}
	writer.Flush() // バッファの内容を実際に書き込む（重要!）
	fmt.Println("output.txt に書き込みました")

	// === 後片付け ===
	os.Remove("sample.txt")
	os.Remove("output.txt")
	fmt.Println("\n一時ファイルを削除しました")
}

// io.Reader を受け取る汎用関数
// ファイルでも文字列でもバッファでも同じように扱える
func printFromReader(label string, r io.Reader) {
	data, _ := io.ReadAll(r)
	fmt.Printf("  [%s] %s\n", label, string(data))
}
