# 008: インターフェース

## このレッスンで学ぶこと

- インターフェースの定義
- **暗黙的実装**（Go最大の特徴の一つ）
- インターフェースを使ったポリモーフィズム
- 空インターフェース `interface{}` / `any`
- 型アサーション

## コード解説

### インターフェースの定義

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

インターフェースは**メソッドの集合**を定義する。

### 暗黙的実装

```go
type Circle struct { Radius float64 }

func (c Circle) Area() float64 { ... }
func (c Circle) Perimeter() float64 { ... }
// Circle は自動的に Shape インターフェースを満たす
```

Java の `implements` キーワードは不要。
メソッドが揃っていれば**自動的に**インターフェースを満たす。

これにより、**後から**既存の型にインターフェースを適用できる。

### ポリモーフィズム

```go
func printInfo(s Shape) {
    fmt.Printf("面積: %.2f\n", s.Area())
}

printInfo(Circle{Radius: 5})    // OK
printInfo(Rectangle{W: 3, H: 4}) // OK
```

### 空インターフェース

```go
var x interface{} // 何でも入る（Go 1.18 以降は any）
x = 42
x = "hello"
```

### 型アサーション

```go
str, ok := x.(string)
if ok {
    fmt.Println("文字列:", str)
}
```

## 設計指針

- **小さなインターフェース**が良い（メソッド1〜2個）
- 標準ライブラリの例: `io.Reader`（`Read` メソッドだけ）, `fmt.Stringer`（`String` メソッドだけ）
- 「使う側」でインターフェースを定義するのが Go流

## やってみよう

1. `Animal` インターフェース（`Speak() string`）を定義し、`Dog` と `Cat` で実装してみよう
2. `fmt.Stringer` を自分の構造体に実装してみよう
3. `io.Reader` を実装する型を作ってみよう
