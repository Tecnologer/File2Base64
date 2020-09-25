package main

import (
	"path/filepath"
	"regexp"
)

var prefixRegex *regexp.Regexp

func init() {
	prefixRegex = regexp.MustCompile(`(?m)^data:.*;base64,`)
}

func decode() {
	if *outputFile == "" {
		panic("to decode base64 file is necesary the output file")
	}

	content, err := getFileContent()
	if err != nil {
		panic(err)
	}

	ext := filepath.Ext(*outputFile)
	if prefixRegex.MatchString(string(content)) {
		if ext == "" {
			ext = getFileExt(string(content))
		}
		content = prefixRegex.ReplaceAll(content, []byte{})
	}

	err = writeFile(content)
	if err != nil {
		panic(err)
	}
}

func getFileExt(content string) string {
	return ""
}
