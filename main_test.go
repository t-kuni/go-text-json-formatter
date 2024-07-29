package main

import (
	"embed"
	_ "embed"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"os"
	"testing"
)

//go:embed testCases/*
var testCases embed.FS

func TestBeautify(t *testing.T) {
	for i := 1; ; i++ {
		inputFilePath := fmt.Sprintf("testCases/input_%d.go", i)
		expectFilePath := fmt.Sprintf("testCases/expect_%d.go", i)
		input, err := fs.ReadFile(testCases, inputFilePath)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				break
			}
		}
		expect, err := fs.ReadFile(testCases, expectFilePath)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				break
			}
		}

		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			output := beautify(string(input))
			assert.Equal(t, string(expect), output)
		})
	}
}
