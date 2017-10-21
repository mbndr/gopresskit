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
	// delete previous outputs if forced
	if p.ForceMode {
		os.RemoveAll(p.OutputPath)
	}

	// setup base folder structure
	err := p.setupOutputFolder(p.OutputPath, []string{"zip", "css"})
	if err != nil {
		return err
	}

	// add file to tell gh-pages not to use jekyll
	// _header, _icon, _logo not rendered there if not generated
	err = writeFile(join(p.OutputPath, ".nojekyll"), []byte{})
	if err != nil {
		return err
	}

	// process company
	c, err := p.processCompany()
	if err != nil {
		return err
	}

	// process games
	files, err := ioutil.ReadDir(join(p.InputPath, "games"))
	if err != nil {
		return err
	}
	for _, f := range files {
		err := p.processGame(f, *c)
		if err != nil {
			return err
		}
	}

	// generate static files (css etc)
	err = p.generateStaticFiles()
	if err != nil {
		return err
	}

	return nil
}

// processCompany processes the company level
func (p Presskit) processCompany() (*company, error) {
	// copy media directories
	for _, d := range []string{"images", "videos"} {
		copyDir(
			join(p.InputPath, d),
			join(p.OutputPath, d),
		)
	}

	// parse data file
	fileData, err := ioutil.ReadFile(join(p.InputPath, "company."+p.Parser.Extension()))
	if err != nil {
		return nil, err
	}
	company, err := p.Parser.Company(fileData)
	if err != nil {
		return nil, err
	}

	// used medias
	media, err := newMedia(p.OutputPath, *company)
	if err != nil {
		return nil, err
	}

	// generate image zip files
	err = media.generateZip(p.OutputPath)
	if err != nil {
		return nil, err
	}

	// get a list of game titles to compile into company template
	// without this processCompany and processGame couldn't be outsourced
	// TODO not optimal
	gameList := p.gameList()

	// render html
	html, err := renderCompany(*company, *media, gameList)
	if err != nil {
		return nil, err
	}
	err = writeFile(join(p.OutputPath, "index.html"), html)
	if err != nil {
		return nil, err
	}

	return company, nil
}

// processGame processes a game level
func (p Presskit) processGame(f os.FileInfo, c company) error {
	// path of the game data
	gameInputPath := join(p.InputPath, "games", f.Name())
	gameOutputPath := join(p.OutputPath, "games", f.Name())

	// setup base folder structure
	err := p.setupOutputFolder(gameOutputPath, []string{"zip"})
	if err != nil {
		return err
	}

	// copy media directories
	for _, d := range []string{"images", "videos"} {
		copyDir(
			join(gameInputPath, d),
			join(gameOutputPath, d),
		)
	}

	// parse data file
	fileData, err := ioutil.ReadFile(join(gameInputPath, "game."+p.Parser.Extension()))
	if err != nil {
		return err
	}
	game, err := p.Parser.Game(fileData)
	if err != nil {
		return err
	}

	// used medias
	media, err := newMedia(gameOutputPath, *game)
	if err != nil {
		return err
	}

	// generate image zip files
	err = media.generateZip(gameOutputPath)
	if err != nil {
		return err
	}

	// render html
	html, err := renderGame(c, *game, *media)
	if err != nil {
		return err
	}
	err = writeFile(join(gameOutputPath, "index.html"), html)
	if err != nil {
		return err
	}

	return nil
}

// setupOutputFolder creates the output folder and a few subdirs.
// used by company and games
func (p Presskit) setupOutputFolder(outputPath string, folders []string) error {
	// don't override
	if _, err := os.Stat(outputPath); !os.IsNotExist(err) {
		return errors.New("output path already exists")
	}

	// create necessary subpaths
	for _, f := range folders {
		err := os.MkdirAll(join(outputPath, f), 0700)
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

// get a list of only the game titles (only temporary)
// TODO this file reading is way to redundant to above!
func (p Presskit) gameList() map[string]string {
	list := make(map[string]string)

	files, _ := ioutil.ReadDir(join(p.InputPath, "games"))
	for _, f := range files {
		// parse data file
		fileData, _ := ioutil.ReadFile(join(p.InputPath, "games", f.Name(), "game."+p.Parser.Extension()))
		game, _ := p.Parser.Game(fileData)
		list[f.Name()] = game.Title
	}

	return list
}
