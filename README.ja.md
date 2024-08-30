lang: [EN](README.md) [JA](README.ja.md)

# 概要

これはテキストリテラル（`` ` ``で囲まれたもの）内のJSON文字列を整形するCLIツールです。

整形前：

```go
	json = `

           {
  "key1":         "value1"      ,
        "key2": {
    "key3": 
        "value3"
  }
}

     `
```

整形後：

```go
	json = `
{
  "key1": "value1",
  "key2": {
    "key3": "value3"
  }
}`
```

## インストール

```bash
go install github.com/t-kuni/go-text-json-formatter@X.X
```

`X.X`は使用しているGoのバージョンに置き換えてください。

## 更新

```bash
go install -a github.com/t-kuni/go-text-json-formatter@X.X
```

## 使用方法

```bash
# ファイルを指定した場合、単一のファイルをフォーマットします
go-text-json-formatter ./path/to/file.go
# ディレクトリを指定した場合、ディレクトリ内を再帰的にフォーマットします
go-text-json-formatter ./path/to/dir
```

このコマンドは指定されたディレクトリ下のGoファイルにあるテキストリテラル内のJSON文字列を整形します。