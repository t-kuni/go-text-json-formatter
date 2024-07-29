package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	input, err := os.Open("example/example.go")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	code, err := io.ReadAll(input)
	if err != nil {
		panic(err)
	}

	output, err := os.Create("example/result.go")
	if err != nil {
		panic(err)
	}
	defer output.Close()

	formattedCode := beautify(string(code))

	_, err = output.WriteString(formattedCode)
	if err != nil {
		panic(err)
	}
}

const ignoreJsonLength = 20

func beautify(code string) string {
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

				var obj interface{}
				err := json.Unmarshal([]byte(textBody), &obj)
				if err != nil {
					fmt.Fprintln(os.Stderr, "[WARN] ", err)
					return true
				}

				prettyJSON, err := json.MarshalIndent(obj, "", "  ")
				if err != nil {
					fmt.Fprintln(os.Stderr, "[WARN] ", err)
					return true
				}

				lit.Value = "`\n" + string(prettyJSON) + "`\n\n"
			}
		}
		return true
	})

	var buf bytes.Buffer
	err = printer.Fprint(&buf, fset, f)
	if err != nil {
		log.Fatal(err)
	}

	return buf.String()
}
