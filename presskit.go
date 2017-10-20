package presskit

import (
	"errors"
	"io/ioutil"
	"os"
	//"github.com/sanity-io/litter"
)

// Presskit is created by the caller script and wraps all generation functions
type Presskit struct {
	// parser manages the parsing from the data file (e.g. company.xml) to a company object
	Parser Parser

	// input paths contains the company data, games and assets
	InputPath string

	// output path contains all generated data
	OutputPath string

	// if enabled, the output folder will removed and overwritten instead of a error
	ForceMode bool
}

// Generate handles the whole generation process
func (p Presskit) Generate() error {
	// create output folder
	err := p.setupOutputFolder()
	if err != nil {
		return err
	}

	// process company
	err = p.processCompany()
	if err != nil {
		return err
	}

	/* TODO
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
	}
	*/

	// generate static files (css etc)
	err = p.generateStaticFiles()
	if err != nil {
		return err
	}

	return nil
}

// processCompany processes the company level
func (p Presskit) processCompany() error {
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
	fileData, err := ioutil.ReadFile(join(p.InputPath, "company."+p.Parser.Extension()))
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
	html, err := renderCompany(*company, *media)
	if err != nil {
		return err
	}
	err = writeFile(join(p.OutputPath, "index.html"), html)
	if err != nil {
		return err
	}

	return nil
}

/* TODO
// processGame processes a game level
func (p Presskit) processGame(f os.FileInfo) error {

}
*/

// setupOutputFolder creates the output folder and a few subdirs
func (p Presskit) setupOutputFolder() error {
	if p.ForceMode {
		os.RemoveAll(p.OutputPath)
	}

	if _, err := os.Stat(p.OutputPath); !os.IsNotExist(err) && !p.ForceMode {
		return errors.New("output path already exists")
	}

	// create necessary subpaths
	folders := []string{"css", "zip"}

	for _, f := range folders {
		err := os.MkdirAll(join(p.OutputPath, f), 0700)
		if err != nil {
			return err
		}
	}

	// create file for disabling jekyll on gh-pages
	err := writeFile(join(p.OutputPath, ".nojekyll"), []byte{})
	if err != nil {
		return err
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
