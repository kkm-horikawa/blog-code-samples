# === ビルドステージ ===
FROM golang:1.23 AS builder

WORKDIR /app

# 依存ファイルを先にコピー（キャッシュ効率化）
COPY go.mod ./
RUN go mod download

# ソースコードをコピーしてビルド
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/server main.go

# === 実行ステージ ===
# scratch = 完全に空のイメージ（OS すらない）
FROM scratch

# ビルドしたバイナリだけコピー
COPY --from=builder /app/server /server

# ポート公開
EXPOSE 8080

# 実行
ENTRYPOINT ["/server"]
