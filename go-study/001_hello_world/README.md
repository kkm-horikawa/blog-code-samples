# 001: Hello World

## このレッスンで学ぶこと

- Go プログラムの基本構造
- `package main` と `func main()` の意味
- `fmt.Println` で文字列を出力する方法

## コード解説

```go
package main
```

すべてのGoファイルは `package` 宣言から始まる。
`main` パッケージは特別で、**実行可能なプログラムのエントリポイント**になる。

```go
import "fmt"
```

`fmt` は標準ライブラリのパッケージ。フォーマット付き入出力（`Println`, `Printf`, `Sprintf` 等）を提供する。

```go
func main() {
    fmt.Println("Hello world!!")
}
```

- `func main()` がプログラムの開始地点。引数も戻り値もない
- `fmt.Println` は引数を出力して改行する

## 実行方法

```bash
go run main.go
```

## やってみよう

1. `"Hello world!!"` を自分の名前に変えて実行してみよう
2. `fmt.Println` を2行に増やして、2行出力してみよう
3. `fmt.Printf("名前: %s, 年齢: %d\n", "太郎", 25)` を試してみよう
