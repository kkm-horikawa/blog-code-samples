# 021: Docker とデプロイ

## このレッスンで学ぶこと

- Go のビルドとシングルバイナリ
- マルチステージ Dockerfile
- 本番向けビルドオプション
- Makefile でタスクをまとめる

## コード解説

### Go のビルドの強み

```bash
go build -o myapp main.go
```

これだけで**単一の実行可能バイナリ**が生成される。

- ランタイム不要（Python/Node のようにインタプリタが要らない）
- 依存ライブラリも全部バイナリに含まれる
- コンテナイメージが**極めて小さく**できる

### マルチステージ Dockerfile

```dockerfile
# ビルドステージ
FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/server ./main.go

# 実行ステージ（scratch = 空のイメージ）
FROM scratch
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]
```

| Django (yorozu) | Go |
|----------------|-----|
| `python:3.12-slim` ベース (~150MB) | `scratch` ベース (~10MB) |
| pip install + 仮想環境 | go build のみ |
| uvicorn でサーバー起動 | バイナリ直接実行 |
| requirements.txt / uv.lock | go.mod / go.sum |

### ビルドオプション

```bash
# 静的リンク（外部ライブラリに依存しない）
CGO_ENABLED=0 go build -o myapp

# サイズ最適化（デバッグ情報を削除）
go build -ldflags="-s -w" -o myapp

# クロスコンパイル（Mac で Linux バイナリを生成）
GOOS=linux GOARCH=amd64 go build -o myapp
GOOS=linux GOARCH=arm64 go build -o myapp  # ARM (Graviton等)
```

### Makefile

```makefile
.PHONY: build run test clean docker

build:
    go build -o bin/server ./cmd/server

run:
    go run ./cmd/server

test:
    go test -v -cover ./...

clean:
    rm -rf bin/

docker:
    docker build -t myapp .
```

Django プロジェクトの Makefile と同じ感覚で使える。

## デプロイ先の選択肢

| 方式 | 特徴 | yorozu で使っている対応先 |
|------|------|------------------------|
| **バイナリ直接** | EC2/VM にバイナリを置くだけ | - |
| **Docker + ECS** | コンテナ実行 | ECS Fargate / EC2 |
| **Docker + App Runner** | フルマネージド | App Runner |
| **Docker + Kubernetes** | 大規模向け | EKS |
| **Lambda** | サーバーレス | workers/ |

Go はシングルバイナリなので、どの方式でもデプロイが Django より単純になる。

## 実行方法

```bash
# ローカルビルド & 実行
go build -o server main.go
./server

# Docker ビルド & 実行
docker build -t go-study-server .
docker run -p 8080:8080 go-study-server

# curl で確認
curl http://localhost:8080/health
```

## やってみよう

1. `docker build` してイメージサイズを確認してみよう（`docker images`）
2. `GOOS=linux GOARCH=amd64` でクロスコンパイルしてみよう
3. Makefile を作って `make build` / `make test` / `make docker` を定義してみよう
