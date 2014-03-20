package main

import (
	"github.com/wingyplus/lablr/cli"
	"io/ioutil"
	"os"
)

func main() {
	content, _ := ioutil.ReadFile(os.Args[1])
	lablrCLI := cli.LablrCLI{
		Options: []string{},
		Content: content,
	}
	outputContent, _ := lablrCLI.Process()
	ioutil.WriteFile("./iss_2_result.xml", outputContent, os.ModePerm)
}
