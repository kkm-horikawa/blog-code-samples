# Go 学習ノート

## ディレクトリ構成

各ディレクトリに README.md（解説）と main.go（サンプルコード）が入っている。
VSCode で左右に開いて、README を読みながらコードを書き換えて `go run main.go` で実行する。

| # | ディレクトリ | テーマ |
|---|---|---|
| 001 | hello_world | package / import / Println |
| 002 | variables | var, :=, const, ゼロ値 |
| 003 | functions | 複数戻り値, 可変長引数, 高階関数 |
| 004 | if_for | if初期化文, for, switch |
| 005 | slice | append, スライシング, sort, copy |
| 006 | map | カンマOKイディオム, 走査 |
| 007 | struct | 値/ポインタレシーバ, 埋め込み |
| 008 | interface | 暗黙的実装, 型アサーション |
| 009 | error_handling | errors.Is/As, defer, panic/recover |
| 010 | goroutine | chan, select, WaitGroup, ワーカープール |
| 011 | package | 公開/非公開, サブパッケージ |

---

## Go の定番ライブラリ一覧

### Web フレームワーク / ルーター

| ライブラリ | 位置づけ |
|---|---|
| **net/http** (標準) | Go 1.22 でルーティング強化。まずこれから |
| **gin** | 最も人気のWebフレームワーク。ミドルウェア・バリデーション付き |
| **echo** | gin と双璧。ドキュメントが良い |
| **chi** | 標準 net/http 互換のルーター。middleware チェーンが使いやすい |

王道: 標準 net/http → gin or chi

### DB / ORM

| ライブラリ | 位置づけ |
|---|---|
| **database/sql** (標準) | 生SQL実行 |
| **sqlx** | database/sql の拡張。構造体への自動マッピング |
| **GORM** | Go最大のORM。マイグレーションもある |
| **sqlc** | SQLからGoコード自動生成。型安全。人気急上昇中 |
| **pgx** | PostgreSQL専用ドライバ（高性能） |

王道: GORM（ORM派）か sqlc（SQL派）。ドライバは pgx

### マイグレーション

| ライブラリ | 位置づけ |
|---|---|
| **golang-migrate/migrate** | 最も定番。SQLファイルで管理 |
| **goose** | シンプル。SQLファイルでマイグレーション管理 |
| **atlas** | スキーマ差分から自動生成 |

### 設定・環境変数

| ライブラリ | 位置づけ |
|---|---|
| **viper** | 最も有名。YAML/JSON/TOML/環境変数を統一的に扱う |
| **envconfig** | 構造体タグで環境変数をバインド。シンプル |
| **godotenv** | .env ファイル読み込み |

### ログ

| ライブラリ | 位置づけ |
|---|---|
| **log/slog** (標準, Go 1.21+) | 構造化ログ。まずこれでOK |
| **zap** (Uber製) | 高性能構造化ログ。大規模プロダクション向け |
| **zerolog** | zap より設定がシンプル。ゼロアロケーション |

王道: まず slog。パフォーマンス要件が出たら zap or zerolog

### テスト

| ライブラリ | 位置づけ |
|---|---|
| **testing** (標準) | フレームワークなしで書く。`go test` で実行 |
| **testify** | アサーション + モック。assert が書きやすくなる |
| **gomock** (uber版) | インターフェースからモック自動生成 |
| **httptest** (標準) | HTTPハンドラのテスト用サーバー |
| **testcontainers-go** | Docker でテスト用DB起動。実DB統合テストに |

王道: 標準 testing + testify

### バリデーション

| ライブラリ | 位置づけ |
|---|---|
| **go-playground/validator** | 構造体タグでバリデーション定義 |

### HTTP クライアント

| ライブラリ | 位置づけ |
|---|---|
| **net/http** (標準) | 標準で十分高機能 |
| **resty** | Python の requests に近い使い勝手 |

### AWS

| ライブラリ | 位置づけ |
|---|---|
| **aws-sdk-go-v2** | boto3 のGo版。公式SDK |

### CLI ツール作成

| ライブラリ | 位置づけ |
|---|---|
| **cobra** | CLI最大手。kubectl, docker, gh コマンドも cobra 製 |
| **urfave/cli** | cobra よりシンプル |

### JSON

| ライブラリ | 位置づけ |
|---|---|
| **encoding/json** (標準) | 基本はこれ |
| **json-iterator** | 標準互換で高速 |
| **sonic** | 最速クラス。大量データ処理向け |

### 並行処理・ユーティリティ

| ライブラリ | 位置づけ |
|---|---|
| **errgroup** (準標準) | ゴルーチンのエラーハンドリング付き並行実行 |
| **sync** (標準) | WaitGroup, Mutex, Once |
| **context** (標準) | キャンセル・タイムアウト伝搬。Go で最重要の概念の一つ |

### API ドキュメント

| ライブラリ | 位置づけ |
|---|---|
| **swaggo/swag** | コメントから Swagger 生成 |
| **oapi-codegen** | OpenAPI スキーマから Go コード生成 |

---

## Web API のスターターセット

```
gin (or chi) + pgx + sqlc + golang-migrate + viper + slog + testify + validator
```

この組み合わせが現在のGoコミュニティで最もメジャーな構成。
