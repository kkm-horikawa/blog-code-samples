# 003: 関数

## このレッスンで学ぶこと

- 関数の定義と呼び出し
- 引数と戻り値
- **複数戻り値**（Goの大きな特徴）
- 名前付き戻り値

## コード解説

### 基本の関数

```go
func add(a int, b int) int {
    return a + b
}
```

`func 関数名(引数) 戻り値の型 { ... }` が基本形。

### 同じ型の引数はまとめられる

```go
func add(a, b int) int {
```

### 複数戻り値

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("0で割れません")
    }
    return a / b, nil
}
```

Go では関数が**複数の値を返せる**。
特に `(結果, error)` のパターンは Go で最も重要なイディオム。

### 名前付き戻り値

```go
func rect(w, h float64) (area, perimeter float64) {
    area = w * h
    perimeter = 2 * (w + h)
    return // 名前付きなので return だけでOK
}
```

戻り値に名前を付けると、関数内で変数として使える。

## 注意点

- Go に関数のオーバーロード（同名で引数違い）はない
- デフォルト引数もない
- 可変長引数は `...` で可能: `func sum(nums ...int) int`

## やってみよう

1. 2つの文字列を結合して返す関数を作ってみよう
2. BMIを計算する関数（身長cm, 体重kg → BMI値）を作ってみよう
3. 最大値と最小値を同時に返す関数を作ってみよう
