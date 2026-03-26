# 004: 条件分岐とループ

## このレッスンで学ぶこと

- `if` / `else if` / `else`
- `if` の初期化文（Go独自）
- `for` ループ（Go唯一のループ構文）
- `switch` 文
- `break` と `continue`

## コード解説

### if 文

```go
if age >= 20 {
    fmt.Println("成人")
} else {
    fmt.Println("未成年")
}
```

カッコ `()` は不要（付けるとエラー）。`{` は同じ行に書く。

### if の初期化文

```go
if score := calcScore(); score >= 80 {
    fmt.Println("合格")
}
// score はこのブロックの外では使えない
```

`if` の中で変数を宣言できる。スコープが限定されるので便利。

### for ループ

Go にはループ構文が **`for` しかない**（`while` がない）。

```go
// 通常の for
for i := 0; i < 5; i++ { }

// while のように使う
for count > 0 { }

// 無限ループ
for { }

// range でスライスを回す
for i, v := range items { }
```

### switch 文

```go
switch day {
case "月", "火", "水", "木", "金":
    fmt.Println("平日")
default:
    fmt.Println("休日")
}
```

Go の switch は**自動で break**する。`fallthrough` で次のcaseに落とせる。

## やってみよう

1. FizzBuzz を書いてみよう（3の倍数→Fizz, 5の倍数→Buzz, 両方→FizzBuzz）
2. 1〜100の合計を `for` で計算してみよう
3. 月(1〜12)を入力して季節を返す `switch` を書いてみよう
