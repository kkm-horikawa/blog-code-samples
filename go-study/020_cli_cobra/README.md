# 020: CLI ツール (cobra)

## このレッスンで学ぶこと

- cobra でCLIツールを作る
- サブコマンドの定義
- フラグ（引数）の扱い
- cobra の構成パターン

## 前提

```bash
go mod init study/020_cli_cobra
go get github.com/spf13/cobra
```

## コード解説

### cobra とは

Go で最も有名な CLI フレームワーク。以下の有名ツールが cobra 製：
- `kubectl`（Kubernetes）
- `docker`
- `gh`（GitHub CLI）
- `hugo`（静的サイトジェネレーター）

### 基本構造

```go
var rootCmd = &cobra.Command{
    Use:   "mytool",
    Short: "ツールの説明",
    Run: func(cmd *cobra.Command, args []string) {
        // 処理
    },
}

func main() {
    rootCmd.Execute()
}
```

### サブコマンド

```
mytool greet --name 太郎
mytool version
mytool user list
mytool user create --name 太郎 --email taro@example.com
```

```go
rootCmd.AddCommand(greetCmd)
rootCmd.AddCommand(versionCmd)
userCmd.AddCommand(userListCmd)
userCmd.AddCommand(userCreateCmd)
rootCmd.AddCommand(userCmd)
```

### フラグの種類

```go
// 永続フラグ: このコマンド + 全サブコマンドで有効
cmd.PersistentFlags().BoolP("verbose", "v", false, "詳細出力")

// ローカルフラグ: このコマンドのみ
cmd.Flags().StringP("name", "n", "", "名前")

// 必須フラグ
cmd.MarkFlagRequired("name")
```

### Django の manage.py との対比

```bash
# Django
python manage.py migrate
python manage.py createsuperuser
python manage.py runserver --port 8080

# cobra
mytool db migrate
mytool user create
mytool server start --port 8080
```

構造は同じ。cobra の方が自由度が高く、自動でヘルプが生成される。

## 実行方法

```bash
go run main.go --help
go run main.go greet --name 太郎
go run main.go greet --name 太郎 --greeting おはよう
go run main.go user list
go run main.go user create --name 花子 --email hanako@example.com
go run main.go version
```

## やってみよう

1. `calc add 3 5` で足し算するサブコマンドを追加してみよう
2. `--output json` フラグでJSON出力に切り替えてみよう
3. `cobra-cli` ジェネレーターでプロジェクトを生成してみよう
