package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <relative_path>")
	}

	relativePath := os.Args[1]
	absolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		panic(err)
	}

	fileInfo, err := os.Stat(absolutePath)
	if err != nil {
		panic(err)
	}

	if fileInfo.IsDir() {
		err = filepath.WalkDir(absolutePath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() && filepath.Ext(path) == ".go" {
				err = processFile(path)
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			log.Fatal("Error walking through the directory:", err)
		}
	} else {
		if filepath.Ext(absolutePath) != ".go" {
			log.Fatal("The file is not a go file.")
		}

		err = processFile(absolutePath)
		if err != nil {
			panic(err)
		}
	}
}

func processFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	code, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	formattedCode, err := beautify(string(code))
	if err != nil {
		return err
	}

	err = os.WriteFile(path, []byte(formattedCode), fileInfo.Mode())
	if err != nil {
		return err
	}

	return nil
}

const ignoreJsonLength = 20

func beautify(code string) (string, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", code, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		if lit, ok := n.(*ast.BasicLit); ok && lit.Kind == token.STRING {
			text := lit.Value
			if text[0] == '`' && text[len(text)-1] == '`' {
				textBody := text[1 : len(text)-1]

				if utf8.RuneCountInString(textBody) <= ignoreJsonLength {
					return true
				}

				var buf bytes.Buffer
				encoder := json.NewEncoder(&buf)
				encoder.SetEscapeHTML(false)
				encoder.SetIndent("", "  ")

				jsonRaw := json.RawMessage(textBody)
				err := encoder.Encode(jsonRaw)
				if err != nil {
					fmt.Fprintln(os.Stderr, "[WARN] ", err)
					return true
				}

				jsonText := strings.TrimSpace(buf.String())

				lit.Value = fmt.Sprintf("`\n%s`\n\n", jsonText)
			}
		}
		return true
	})

	var buf bytes.Buffer
	err = printer.Fprint(&buf, fset, f)
	if err != nil {
		return "", err
	}

	// 謎の空行が入ることがあるのでgo fmtで整形する
	formattedCode, err := format.Source(buf.Bytes())
	if err != nil {
		return "", err
	}

	return string(formattedCode), nil
}
