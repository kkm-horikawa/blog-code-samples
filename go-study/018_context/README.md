# 018: context

## このレッスンで学ぶこと

- `context.Context` とは何か
- タイムアウト（`context.WithTimeout`）
- キャンセル（`context.WithCancel`）
- 値の伝搬（`context.WithValue`）
- HTTP ハンドラでの context の使い方

## コード解説

### context とは

Go で**処理のキャンセル・タイムアウト・期限を伝搬する仕組み**。
Django にはない概念で、Go では最も重要なパターンの一つ。

すべての「時間がかかる処理」は context を第一引数に受け取る慣習：

```go
func FetchUser(ctx context.Context, id int) (*User, error) {
    // ctx がキャンセルされたら処理を中断
}
```

### なぜ必要か

1. **HTTPリクエストのタイムアウト**: クライアントが切断したら処理を止めたい
2. **DB クエリのタイムアウト**: 遅いクエリを一定時間で打ち切りたい
3. **複数ゴルーチンのキャンセル**: 1つ失敗したら全部止めたい

### context.WithTimeout

```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel() // 必ず呼ぶ（リソースリーク防止）

result, err := slowOperation(ctx)
// 3秒以内に終わらなければ ctx.Done() が閉じる
```

### context.WithCancel

```go
ctx, cancel := context.WithCancel(context.Background())

go func() {
    // 何かの条件で
    cancel() // 明示的にキャンセル
}()

<-ctx.Done() // キャンセルされるまで待つ
```

### ctx.Done() の使い方

```go
select {
case <-ctx.Done():
    return ctx.Err() // context.Canceled or context.DeadlineExceeded
case result := <-ch:
    return result
}
```

## ルール

1. **第一引数に `ctx context.Context`** を置く（Go の慣習）
2. **構造体のフィールドに context を保存しない**
3. **`defer cancel()`** を必ず呼ぶ
4. **`context.Background()`** はルート（main やテストの起点）でのみ使う
5. **`context.TODO()`** は「後で適切な context を入れる」プレースホルダ

## Django との対比

Django ではリクエストのタイムアウトは Web サーバー（gunicorn/uvicorn）が管理するので、
アプリケーションコードで意識しない。Go ではアプリケーションが自分で管理する。

| Django | Go |
|--------|-----|
| gunicorn のタイムアウト設定 | `context.WithTimeout` |
| シグナルハンドラ | `context.WithCancel` |
| `request` オブジェクト | `r.Context()` で取得 |

## やってみよう

1. 5秒かかる処理に3秒タイムアウトを設定して、タイムアウトさせてみよう
2. 複数のゴルーチンを `context.WithCancel` で一斉にキャンセルしてみよう
3. HTTP サーバーで `r.Context()` を使って、クライアント切断を検知してみよう
