package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	// inputFile     = flag.String("input", "", "Path of the file")
	minversion    string
	inputFile     string
	outputFile    = flag.String("output", "", "Path of the output file")
	outputConsole = flag.Bool("console", false, "Flag to indicate if the result will be also printed on the console when the output file is specified")
	withType      = flag.Bool("with-type", true, "Flag to indicate if the result will include the type of the file. I.e.: data:image/png;base64,<enconded>")
)

func main() {

	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}

	if inputFile == "" {
		fmt.Println("input file is required. file2Base64 <file_path>")
		flag.PrintDefaults()
		return
	}

	flag.Parse()

	f, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)

	content, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	contentType := http.DetectContentType(content)
	encoded := base64.StdEncoding.EncodeToString(content)

	if *withType {
		encoded = fmt.Sprintf("data:%s;base64,%s", contentType, encoded)
	}

	errorWritting := false
	if *outputFile != "" {
		fOut, err := os.Create(*outputFile)
		if err != nil {
			logrus.Errorf("trying create the result file. Error: %v\n", err)
			errorWritting = true
			return
		}
		defer fOut.Close()
		_, err = fOut.Write([]byte(encoded))
		if err != nil {
			logrus.Errorf("trying to write the result. Error: %v\n", err)
			errorWritting = true
			return
		}
	}

	if *outputFile == "" || errorWritting || *outputConsole {
		fmt.Println(encoded)
	}
}
