# 009: エラーハンドリング

## このレッスンで学ぶこと

- Go のエラー処理の基本（`error` インターフェース）
- `fmt.Errorf` でエラーを作る
- カスタムエラー型
- `errors.Is` / `errors.As`（Go 1.13+）
- エラーのラップ（`%w`）
- `defer` / `panic` / `recover`

## コード解説

### Go にtry/catchはない

Go は例外機構を持たない。代わりに**戻り値で `error` を返す**。

```go
result, err := doSomething()
if err != nil {
    // エラー処理
    return err
}
// 正常処理
```

この `if err != nil` パターンが Go コードの大部分を占める。

### error インターフェース

```go
type error interface {
    Error() string
}
```

`error` は1つのメソッドだけ持つインターフェース。

### エラーの作り方

```go
// 簡単なエラー
errors.New("something went wrong")

// フォーマット付き
fmt.Errorf("ユーザーID %d が見つかりません", id)

// エラーのラップ（元のエラーを保持）
fmt.Errorf("DB操作に失敗: %w", err)
```

### defer

```go
func readFile() {
    f, _ := os.Open("file.txt")
    defer f.Close() // 関数終了時に必ず実行される
    // ファイル操作...
}
```

`defer` はリソース解放に使う。LIFO（後入先出）順で実行される。

## 設計指針

- エラーは握りつぶさない。処理するか、呼び出し元に返す
- `%w` でラップして文脈を追加しながら上位に伝搬させる
- `panic` は本当に回復不能な場面だけ（通常のエラーには使わない）

## やってみよう

1. 年齢を受け取り、0未満や150超ならエラーを返す関数を作ってみよう
2. カスタムエラー型 `ValidationError` を作ってみよう
3. `errors.Is` でエラーの種類を判定する処理を書いてみよう
