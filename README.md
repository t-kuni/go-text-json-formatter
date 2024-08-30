lang: [EN](README.md) [JA](README.ja.md)

# Overview

This is a CLI tool for formatting JSON strings within text literals (enclosed by `` ` ``).

Before formatting:

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

After formatting:

```go
	json = `
{
  "key1": "value1",
  "key2": {
    "key3": "value3"
  }
}`
```

## Installation

```bash
go install github.com/t-kuni/go-text-json-formatter@X.X
```

Replace `X.X` with your Go version.

## Update

```bash
go install -a github.com/t-kuni/go-text-json-formatter@X.X
```

## Usage

```bash
# If a file is specified, it formats a single file
go-text-json-formatter ./path/to/file.go
# If a directory is specified, it formats recursively within the directory
go-text-json-formatter ./path/to/dir

```

This command will format the JSON strings within text literals in Go files under the specified directory.