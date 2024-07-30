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
go install -u github.com/t-kuni/go-text-json-formatter@X.X
```

## Usage

```bash
go-text-json-formatter ./path/to/dir
```

This command will format the JSON strings within text literals in Go files under the specified directory.