package main

import (
	"fmt"
	bf "github.com/russross/blackfriday"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		file, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			println("file not found")
			return
		}
		f := bf.MarkdownCommon(file)
		fmt.Fprintf(os.Stdout, string(f))
	} else {
		println("need help? md-converter MARKDOWN_FILENAME. That's all.")
		return
	}
}
