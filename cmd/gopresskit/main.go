package main

import (
	"flag"
	"log"
	"os"

	"github.com/mbndr/gopresskit"
)

func main() {

	var (
		inputPath  = flag.String("input", ".", "Folder with the presskit data")
		outputPath = flag.String("output", "./press", "Folder for the presskit html")
		forceMode  = flag.Bool("force", false, "Force existing output folder to remove")
		usage  = flag.Bool("help", false, "Show command usage")
	)

	flag.Parse()

	if *usage {
		flag.PrintDefaults()
		os.Exit(1)
	}

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
