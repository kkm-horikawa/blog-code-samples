# 019: ジェネリクス (Go 1.18+)

## このレッスンで学ぶこと

- 型パラメータの基本構文
- 型制約（constraints）
- ジェネリックな関数・構造体
- `comparable` / `any` / カスタム制約
- いつジェネリクスを使うべきか

## コード解説

### なぜジェネリクスが必要か

Go 1.18 以前は、型ごとに同じ関数を書く必要があった：

```go
func MaxInt(a, b int) int { ... }
func MaxFloat(a, b float64) float64 { ... }
func MaxString(a, b string) string { ... }
```

ジェネリクスなら1つで済む：

```go
func Max[T constraints.Ordered](a, b T) T {
    if a > b { return a }
    return b
}
```

### 基本構文

```go
func 関数名[T 制約](引数 T) T {
    // T は型パラメータ
}
```

`[T 制約]` が型パラメータ。呼び出し時に具体的な型が決まる。

### 型制約

```go
// any: 何でもOK（= interface{}）
func Print[T any](v T) { fmt.Println(v) }

// comparable: == で比較可能な型
func Contains[T comparable](slice []T, target T) bool { ... }

// constraints.Ordered: 順序比較可能（<, >, <=, >=）
func Max[T constraints.Ordered](a, b T) T { ... }

// カスタム制約
type Number interface {
    int | int64 | float64
}
func Sum[T Number](nums []T) T { ... }
```

### ジェネリックな構造体

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}
```

### Python / Django との対比

Python は動的型付けなのでジェネリクスが不要（何でも入る）。
Go は静的型付けなので、型安全に汎用コードを書くためにジェネリクスが必要。

```python
# Python: 型を気にしない
def max_value(a, b):
    return a if a > b else b
```

```go
// Go: 型パラメータで型安全に
func Max[T constraints.Ordered](a, b T) T {
    if a > b { return a }
    return b
}
```

## いつ使うべきか

| 使うべき場面 | 使わなくてよい場面 |
|---|---|
| スライス/マップの汎用操作 | 特定の型だけ扱う関数 |
| データ構造（Stack, Queue等） | ビジネスロジック |
| ユーティリティ関数 | 最初の実装（まず具体型で書く） |

**原則**: まず具体的な型で書く → 重複が出たらジェネリクスに抽出する

## やってみよう

1. ジェネリックな `Filter[T any]` 関数を書いてみよう
2. ジェネリックな `Map[T, U any]` 関数を書いてみよう
3. ジェネリックな `Queue[T any]` 構造体を作ってみよう
