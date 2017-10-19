package main

import (
	"flag"
	"os"

	"github.com/mbndr/gopresskit"
	"github.com/mbndr/logo"
)

func main() {

	var (
		log = logo.NewSimpleLogger(os.Stderr, logo.INFO, "", true)

		inputPath  = flag.String("input", ".", "Folder with the presskit data")
		outputPath = flag.String("output", "./press", "Folder for the presskit html")
	)

	flag.Parse()

	p := presskit.Presskit{
		InputPath:  *inputPath,
		OutputPath: *outputPath,
		Parser:     presskit.XMLParser{},
	}

	err := p.Generate()
	if err != nil {
		p.Clean()
		log.Fatal(err)
	}
}
