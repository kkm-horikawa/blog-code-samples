package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// === ルートコマンド ===
var rootCmd = &cobra.Command{
	Use:   "mytool",
	Short: "Go学習用CLIツール",
	Long:  "cobra で作る CLI ツールのサンプル。サブコマンド、フラグ、引数の使い方を学ぶ。",
}

// === greet コマンド ===
var greetName string
var greetMessage string

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "挨拶する",
	Example: `  mytool greet --name 太郎
  mytool greet --name 太郎 --greeting おはよう`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s、%sさん!\n", greetMessage, greetName)
	},
}

// === version コマンド ===
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "バージョンを表示",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mytool v0.1.0")
	},
}

// === user コマンド（親） ===
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "ユーザー管理",
}

// === user list コマンド ===
var outputFormat string

var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "ユーザー一覧を表示",
	Run: func(cmd *cobra.Command, args []string) {
		users := []map[string]string{
			{"name": "太郎", "email": "taro@example.com"},
			{"name": "花子", "email": "hanako@example.com"},
		}

		if outputFormat == "json" {
			data, _ := json.MarshalIndent(users, "", "  ")
			fmt.Println(string(data))
		} else {
			fmt.Println("名前\t\tメール")
			fmt.Println("----\t\t-----")
			for _, u := range users {
				fmt.Printf("%s\t\t%s\n", u["name"], u["email"])
			}
		}
	},
}

// === user create コマンド ===
var createName string
var createEmail string

var userCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "ユーザーを作成",
	RunE: func(cmd *cobra.Command, args []string) error {
		// RunE: エラーを返せる版
		if createName == "" {
			return fmt.Errorf("--name は必須です")
		}
		if createEmail == "" {
			return fmt.Errorf("--email は必須です")
		}

		fmt.Printf("ユーザーを作成しました:\n")
		fmt.Printf("  名前: %s\n", createName)
		fmt.Printf("  メール: %s\n", createEmail)
		return nil
	},
}

// === 永続フラグ（全コマンド共通） ===
var verbose bool

func init() {
	// 永続フラグ（全コマンドで有効）
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "詳細出力")

	// greet コマンドのフラグ
	greetCmd.Flags().StringVarP(&greetName, "name", "n", "世界", "挨拶する相手の名前")
	greetCmd.Flags().StringVarP(&greetMessage, "greeting", "g", "こんにちは", "挨拶メッセージ")

	// user list コマンドのフラグ
	userListCmd.Flags().StringVarP(&outputFormat, "output", "o", "table", "出力形式 (table|json)")

	// user create コマンドのフラグ（必須）
	userCreateCmd.Flags().StringVarP(&createName, "name", "n", "", "ユーザー名")
	userCreateCmd.Flags().StringVarP(&createEmail, "email", "e", "", "メールアドレス")
	userCreateCmd.MarkFlagRequired("name")
	userCreateCmd.MarkFlagRequired("email")

	// サブコマンドの登録
	rootCmd.AddCommand(greetCmd)
	rootCmd.AddCommand(versionCmd)

	userCmd.AddCommand(userListCmd)
	userCmd.AddCommand(userCreateCmd)
	rootCmd.AddCommand(userCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
