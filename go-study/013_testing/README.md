# 013: テスト

## このレッスンで学ぶこと

- `go test` の使い方
- テスト関数の命名規則
- **テーブル駆動テスト**（Go の最重要テストパターン）
- testify によるアサーション
- テストヘルパーとサブテスト
- カバレッジの確認

## コード解説

### テストファイルの規則

```
calculator.go       ← 本体コード
calculator_test.go  ← テスト（_test.go で終わるファイル）
```

- テストファイルは `_test.go` で終わる
- テスト関数は `Test` で始まり `*testing.T` を受け取る
- ビルドには含まれない（テスト時のみコンパイル）

### 基本のテスト

```go
func TestAdd(t *testing.T) {
    got := Add(2, 3)
    want := 5
    if got != want {
        t.Errorf("Add(2, 3) = %d, want %d", got, want)
    }
}
```

### テーブル駆動テスト（Go で最も重要）

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"正の数", 2, 3, 5},
        {"ゼロ", 0, 0, 0},
        {"負の数", -1, -2, -3},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
            }
        })
    }
}
```

テストケースをスライスに並べて `t.Run` で回す。
**Go のテストコードの8割はこのパターン**。

### testify（外部ライブラリ）

```go
import "github.com/stretchr/testify/assert"

func TestAdd(t *testing.T) {
    assert.Equal(t, 5, Add(2, 3))
}
```

pytest の `assert` に近い書き味になる。

## 実行方法

```bash
go test ./...                    # 全テスト実行
go test -v ./...                 # 詳細出力
go test -run TestAdd ./...       # 特定テストだけ実行
go test -cover ./...             # カバレッジ表示
go test -coverprofile=cover.out  # カバレッジファイル出力
go tool cover -html=cover.out    # ブラウザで可視化
```

## pytest との対比

| pytest | Go testing |
|--------|-----------|
| `assert x == 1` | `if got != want { t.Errorf(...) }` |
| `@pytest.mark.parametrize` | テーブル駆動テスト |
| `conftest.py` のフィクスチャ | `TestMain` or ヘルパー関数 |
| `pytest-cov` | `go test -cover`（標準機能） |

## やってみよう

1. `Multiply` 関数を追加してテーブル駆動テストを書いてみよう
2. `go test -cover` でカバレッジを確認してみよう
3. testify の `assert.Error` でエラーケースのテストを書いてみよう
