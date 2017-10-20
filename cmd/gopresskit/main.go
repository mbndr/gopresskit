package main

import (
	"flag"
	"log"

	"github.com/mbndr/gopresskit"
)

func main() {

	var (
		inputPath  = flag.String("input", ".", "Folder with the presskit data")
		outputPath = flag.String("output", "./press", "Folder for the presskit html")
		forceMode  = flag.Bool("force", false, "Force existing output folder to remove")
	)

	flag.Parse()

	p := presskit.Presskit{
		InputPath:  *inputPath,
		OutputPath: *outputPath,
		Parser:     presskit.XMLParser{},
		ForceMode:  *forceMode,
	}

	err := p.Generate()
	if err != nil {
		log.Fatal(err)
	}
}
