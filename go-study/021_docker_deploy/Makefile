.PHONY: build run test clean docker docker-run

# バイナリビルド
build:
	go build -o bin/server main.go

# 開発サーバー起動
run:
	go run main.go

# テスト実行
test:
	go test -v -cover ./...

# ビルド成果物を削除
clean:
	rm -rf bin/

# Docker イメージビルド
docker:
	docker build -t go-study-server .

# Docker コンテナ実行
docker-run:
	docker run -p 8080:8080 go-study-server

# クロスコンパイル（Linux AMD64）
build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/server-linux main.go

# クロスコンパイル（Linux ARM64 / Graviton）
build-arm:
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o bin/server-arm main.go
