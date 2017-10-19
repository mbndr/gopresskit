package presskit

import (
	"errors"
	"os"
	"io/ioutil"

	//"github.com/sanity-io/litter"
)

type Presskit struct {
	Parser     Parser
	InputPath  string
	OutputPath string
}

// Generate generates the static html files
func (p Presskit) Generate() error {
	// process company
	err := p.processCompany()
	if err != nil {
		return err
	}

	/*
	// process games
	files, err := ioutil.ReadDir(filepath.Join(p.InputPath, "games"))
	if err != nil {
		return err
	}
	for _, f := range files {
		err := p.processGame(f)
		if err != nil {
			return err
		}
	}*/

	// generate static files (css etc)
	err = p.generateStaticFiles()
	if err != nil {
		return err
	}

	return nil
}


// processCompany processes the company level
func (p Presskit) processCompany() error {

	// create output folder
	err := p.setupOutputFolder()
	if err != nil {
		return err
	}

	// copy images and videos
	copyDirs := []string{"images", "videos"}
	for _, d := range copyDirs {
		err := copyDir(
			join(p.InputPath, d),
			join(p.OutputPath, d),
		)
		if err != nil {
			return err
		}
	}

	// parse data file
	fileData, err := ioutil.ReadFile(join(p.InputPath, "company." + p.Parser.Extension()))
	if err != nil {
		return err
	}
	company, err := p.Parser.Company(fileData)
	if err != nil {
		return err
	}

	// used medias
	media, err := newMedia(p.OutputPath, *company)
	if err != nil {
		return err
	}

	// generate image zip files
	err = media.generateZip(p.OutputPath)
	if err != nil {
		return err
	}

	// render html
	html, err := renderCompany(*company, *media, p.OutputPath)
	if err != nil {
		return err
	}
	err = writeFile(join(p.OutputPath, "index.html"), html)
	if err != nil {
		return err
	}

	return nil
}

/*
// processCompany processes a game level
func (p Presskit) processGame(f os.FileInfo) error {
	// paths relative to the
	gameInputPath := filepath.Join(p.InputPath, "games", f.Name())
	gameOutputPath := filepath.Join(p.OutputPath, "games", f.Name())
	litter.Dump(gameInputPath)
	litter.Dump(gameOutputPath)

	// create output dir
	err := os.MkdirAll(fullOutputPath, 0700)
	if err != nil {
		return err
	}
	// copy images
	err = copyDir(
		filepath.Join(fullInputPath, "images"),
		filepath.Join(fullOutputPath, "images"),
	)
	if err != nil {
		return err
	}

	// parse data file
	fileData, err := p.readInputFile(filepath.Join("" "game.") + p.Parser.Extension())
	if err != nil {
		return err
	}

	// TODO render html of game page
	// TODO zip images / logo&icon
	return nil
}*/

// setupOutputFolder creates the output folder and a few subdirs
// TODO remove debug removement / add flag
func (p Presskit) setupOutputFolder() error {
	os.RemoveAll(p.OutputPath)

	if _, err := os.Stat(p.OutputPath); !os.IsNotExist(err) {
		return errors.New("output path already exists")
	}

	// create neccessary subpaths
	folders := []string{"css", "zip"}

	for _, f := range folders {
		err := os.MkdirAll(join(p.OutputPath, f), 0700)
		if err != nil {
			return err
		}
	}
	return nil
}

// generateStaticFiles generates all needed files stored binary
// TODO list of files to generate if more files are added
func (p Presskit) generateStaticFiles() error {
	cssFile := "css/style.css"
	style, err := Asset(cssFile)
	if err != nil {
		return err
	}
	err = writeFile(join(p.OutputPath, cssFile), style)
	if err != nil {
		return err
	}
	return nil
}


// Clean removes the output folder and should only be called on error
func (p Presskit) Clean() {
	os.RemoveAll(p.OutputPath)
}
