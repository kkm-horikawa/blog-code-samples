# 011: パッケージとモジュール

## このレッスンで学ぶこと

- `go mod init` でモジュールを作る
- パッケージの分割
- 公開（大文字始まり）と非公開（小文字始まり）
- 外部パッケージのインポート
- プロジェクト構成のベストプラクティス

## コード解説

### モジュールの初期化

```bash
mkdir myapp && cd myapp
go mod init myapp
```

`go.mod` ファイルが生成される。これがモジュールのルート。

### パッケージ分割

```
myapp/
├── go.mod
├── main.go          # package main
└── greeting/
    └── greeting.go  # package greeting
```

ディレクトリ = パッケージ。ディレクトリ名がパッケージ名になる。

### 公開 / 非公開

```go
package greeting

// Hello は公開（大文字始まり）→ 外部から使える
func Hello(name string) string {
    return format(name) // 非公開関数を内部で使う
}

// format は非公開（小文字始まり）→ パッケージ外からは見えない
func format(name string) string {
    return "こんにちは、" + name + "さん!"
}
```

Go の可視性ルールはシンプル: **大文字で始まれば公開、小文字なら非公開**。

### 外部パッケージの利用

```bash
go get github.com/fatih/color
```

```go
import "github.com/fatih/color"
color.Red("エラー!")
```

`go get` でインストール → `import` で使う。依存は `go.mod` に自動記録される。

## プロジェクト構成の例

```
myapp/
├── cmd/
│   └── server/
│       └── main.go     # エントリポイント
├── internal/            # 外部から import 不可
│   ├── handler/
│   └── model/
├── pkg/                 # 外部から import 可
│   └── utils/
├── go.mod
└── go.sum
```

- `cmd/`: 実行可能バイナリのエントリポイント
- `internal/`: プロジェクト内部でのみ使うコード
- `pkg/`: 外部に公開してよいライブラリ

## やってみよう

1. `calculator` パッケージを作り、`Add`, `Subtract` 関数を `main` から呼んでみよう
2. 非公開関数を外部から呼ぼうとするとどうなるか確認しよう
3. `go get` で好きな外部パッケージをインストールして使ってみよう
