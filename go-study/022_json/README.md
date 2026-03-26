# 022: JSON と構造体タグ

## このレッスンで学ぶこと

- `encoding/json` によるJSON変換
- 構造体タグ（`json:"name"`）
- Marshal（Go → JSON）と Unmarshal（JSON → Go）
- `omitempty` / `-` / カスタム型
- ネストした構造体・スライスの扱い
- API レスポンス設計パターン

## コード解説

### 構造体 → JSON（Marshal）

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age,omitempty"` // ゼロ値なら省略
}

user := User{Name: "太郎", Email: "taro@example.com"}
data, _ := json.Marshal(user)
// {"name":"太郎","email":"taro@example.com"}
// age はゼロ値(0)なので omitempty で省略される
```

### JSON → 構造体（Unmarshal）

```go
jsonStr := `{"name":"花子","email":"hanako@example.com","age":25}`
var user User
json.Unmarshal([]byte(jsonStr), &user)
```

### 構造体タグの詳細

```go
type Example struct {
    Name     string `json:"name"`           // キー名を "name" に
    Age      int    `json:"age,omitempty"`   // ゼロ値なら省略
    Password string `json:"-"`              // JSON に含めない
    Score    float64 `json:"score,string"`   // 文字列として出力
}
```

| タグ | 効果 |
|------|------|
| `json:"name"` | JSON キー名を指定 |
| `json:",omitempty"` | ゼロ値（0, "", nil, false）なら省略 |
| `json:"-"` | JSON に含めない |
| `json:",string"` | 数値を文字列として出力 |

### DRF Serializer との対比

| DRF Serializer | Go 構造体タグ |
|---------------|-------------|
| `source="user_name"` | `json:"user_name"` |
| `required=False` | `omitempty` or ポインタ型 |
| `write_only=True` | レスポンス用とリクエスト用で構造体を分ける |
| `SerializerMethodField` | カスタム `MarshalJSON` |

## API 設計パターン

Go の API では、リクエスト用とレスポンス用で構造体を分けるのが一般的：

```go
// リクエスト用（バリデーション付き）
type CreateUserRequest struct {
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
}

// レスポンス用（パスワードなどを含まない）
type UserResponse struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

## やってみよう

1. ネストしたJSONをパースする構造体を定義してみよう
2. `MarshalJSON` をカスタム実装して日付フォーマットを変えてみよう
3. 外部APIのレスポンスJSONに対応する構造体を作ってみよう
