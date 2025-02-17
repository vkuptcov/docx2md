package main

import (
	"flag"
	"fmt"
	"github.com/mattn/docx2md/docx2markdown"
	"log"
	"os"
	"runtime"
)

const name = "docx2md"
const version = "0.0.9"

var revision = "HEAD"

func main() {
	var embed bool
	var showVersion bool
	flag.BoolVar(&embed, "embed", false, "embed resources")
	flag.BoolVar(&showVersion, "v", false, "Print the version")
	flag.Parse()
	if showVersion {
		fmt.Printf("%s %s (rev: %s/%s)\n", name, version, revision, runtime.Version())
		return
	}
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	for _, arg := range flag.Args() {
		if res, err := docx2markdown.Docx2md(arg, embed); err == nil {
			fmt.Print(res)
		} else {
			log.Fatal(err)
		}
	}
}
