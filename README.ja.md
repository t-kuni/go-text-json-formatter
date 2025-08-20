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

## 新バージョン対応

1. masterブランチからバージョンブランチ（`1.24`など）を作成
2. リポジトリ全体の `1.22` を任意のバージョンに置換する
3. 以下のコマンドでテストを実行する（`1.22`の部分を任意のバージョンに置き換えること）
  * `docker run --rm -v "$(pwd)":/workspace -w /workspace golang:1.22 go test -v`

バージョンを置換する以外の修正は masterブランチに対して変更をコミットしたあと、apply_changes.sh を用いること