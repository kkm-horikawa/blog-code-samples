# 023: io.Reader / io.Writer とファイル操作

## このレッスンで学ぶこと

- `io.Reader` / `io.Writer` インターフェース
- ファイルの読み書き
- `os.ReadFile` / `os.WriteFile`（簡易版）
- `bufio` でバッファ付き読み込み
- `io` パッケージのユーティリティ

## コード解説

### io.Reader / io.Writer とは

Go の I/O の根幹となる2つのインターフェース：

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

このインターフェースを実装しているもの：

| 型 | Reader | Writer |
|---|---|---|
| `*os.File` | Yes | Yes |
| `*bytes.Buffer` | Yes | Yes |
| `*strings.Reader` | Yes | - |
| `http.ResponseWriter` | - | Yes |
| `http.Request.Body` | Yes | - |
| `os.Stdout` / `os.Stderr` | - | Yes |

**「ファイル」「ネットワーク」「メモリ」を同じインターフェースで扱える**のが Go の設計思想。

### ファイル操作の方法（3段階）

```go
// 1. 簡易版（小さなファイル向け）
data, _ := os.ReadFile("file.txt")
os.WriteFile("file.txt", data, 0644)

// 2. 標準版（大きなファイル向け）
f, _ := os.Open("file.txt")
defer f.Close()
scanner := bufio.NewScanner(f)

// 3. 低レベル（細かい制御が必要な場合）
f, _ := os.OpenFile("file.txt", os.O_RDWR|os.O_CREATE, 0644)
defer f.Close()
f.Read(buf)
f.Write(data)
```

### Django / Python との対比

```python
# Python
with open("file.txt") as f:
    content = f.read()

for line in open("file.txt"):
    print(line)
```

```go
// Go（簡易版）
data, _ := os.ReadFile("file.txt")

// Go（行ごと）
f, _ := os.Open("file.txt")
defer f.Close()
scanner := bufio.NewScanner(f)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}
```

Python の `with` 文 = Go の `defer f.Close()`

## io パッケージの便利関数

| 関数 | 用途 |
|------|------|
| `io.Copy(dst, src)` | Reader から Writer にコピー |
| `io.ReadAll(r)` | Reader の全内容を読む |
| `io.TeeReader(r, w)` | 読みながら別の Writer にも書く |
| `io.MultiWriter(w1, w2)` | 複数の Writer に同時書き込み |

## やってみよう

1. テキストファイルを読み込んで、行番号付きで出力してみよう
2. `io.MultiWriter` でファイルと標準出力に同時に書いてみよう
3. HTTP レスポンスボディをファイルに保存する処理を書いてみよう
