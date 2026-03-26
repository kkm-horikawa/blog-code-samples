# 017: プロジェクト構成

## このレッスンで学ぶこと

- Go プロジェクトの標準的なディレクトリ構成
- `cmd/` `internal/` `pkg/` の使い分け
- レイヤー分離（ハンドラ → サービス → リポジトリ）
- 依存性の方向

## このレッスンは「読むだけ」

実行可能なコードはない。構成の理解が目的。

## 標準的なプロジェクト構成

```
myapp/
├── cmd/                       # エントリポイント
│   └── server/
│       └── main.go            # func main() はここだけ
├── internal/                  # このプロジェクト専用コード（外部からimport不可）
│   ├── handler/               # HTTPハンドラ（= Django views）
│   │   └── user_handler.go
│   ├── service/               # ビジネスロジック
│   │   └── user_service.go
│   ├── repository/            # DB操作（= Django models のクエリ部分）
│   │   └── user_repository.go
│   ├── model/                 # データモデル（= Django models の定義部分）
│   │   └── user.go
│   └── middleware/            # ミドルウェア
│       └── auth.go
├── pkg/                       # 外部に公開してよい汎用コード
│   └── validator/
├── config/                    # 設定読み込み
│   └── config.go
├── migrations/                # DBマイグレーションファイル
├── go.mod
├── go.sum
├── Makefile
└── Dockerfile
```

## 各ディレクトリの役割

### `cmd/` — エントリポイント

```go
// cmd/server/main.go
func main() {
    cfg := config.Load()
    db := database.Connect(cfg.DatabaseURL)

    userRepo := repository.NewUserRepository(db)
    userSvc := service.NewUserService(userRepo)
    userHandler := handler.NewUserHandler(userSvc)

    r := gin.Default()
    userHandler.RegisterRoutes(r)
    r.Run(cfg.Port)
}
```

`main.go` の責務は**組み立て（ワイヤリング）だけ**。ビジネスロジックを書かない。

### `internal/` — プロジェクト専用コード

Go の特殊ルール: `internal/` 配下は**同じモジュール内からしか import できない**。
外部パッケージが依存するのを防ぐ。

### `pkg/` — 公開可能な汎用コード

他のプロジェクトから `import` されても問題ないもの。
ただし、最近は `pkg/` を使わない流派も多い（全部 `internal/` に入れる）。

## レイヤー構成

```
Handler（リクエスト受け取り・レスポンス返却）
    ↓
Service（ビジネスロジック）
    ↓
Repository（DB操作）
    ↓
Database
```

### Django との対比

| Django | Go |
|--------|-----|
| `views.py` | `handler/` |
| `serializers.py` | `handler/` 内のリクエスト/レスポンス構造体 |
| `models.py` のクエリ | `repository/` |
| `models.py` のモデル定義 | `model/` |
| `services.py`（使う場合） | `service/` |
| `urls.py` | `cmd/` か `handler/` の `RegisterRoutes` |
| `settings.py` | `config/` |
| `middleware/` | `middleware/` |

### 依存の方向（重要）

```
handler → service → repository → model
                                    ↑
handler も service も model を参照する（OK）
repository → handler（NG: 逆方向の依存）
```

上位層が下位層に依存する。逆方向は禁止。
これは Django でも同じだが、Go では `internal/` のパッケージ分離で物理的に強制される。

## 小さなプロジェクトの場合

最初から上記の構成を作る必要はない。

```
myapp/
├── main.go          # 最初は1ファイルで十分
├── handler.go       # 大きくなったら分割
├── model.go
├── go.mod
└── Dockerfile
```

**ファイルが大きくなったら分割する**。最初から完璧な構成を目指さない。

## やってみよう

1. 014 や 015 のコードを `handler/` と `model/` に分割してみよう
2. `internal/` に入れて、外部から import できないことを確認しよう
3. `Makefile` を作って `make run` / `make test` / `make build` を定義してみよう
