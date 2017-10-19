package presskit

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

const imageDir = "images"
const videoDir = "videos"

// media holds all paths to images of a page
type media struct {
	Header string
	Icon   string
	Logo   string
	Images []string // images for the image section
	Videos []video  // videos for the video section
}

// newImages reads an image directory and seperates the images
func newMedia(path string, c company) (*media, error) {
	var m media

	// images
	images, err := ioutil.ReadDir(join(path, imageDir))
	if err != nil {
		return nil, err
	}

	for _, f := range images {
		// use only base name to identify usage
		switch strings.TrimSuffix(f.Name(), filepath.Ext(f.Name())) {
		case "_header":
			m.Header = join(imageDir, f.Name())
		case "_logo":
			m.Logo = join(imageDir, f.Name())
		case "_icon":
			m.Icon = join(imageDir, f.Name())
		default:
			m.Images = append(m.Images, join(imageDir, f.Name()))
		}
	}

	// copy external videos
	m.Videos = c.Videos

	// add local videos
	videos, err := ioutil.ReadDir(join(path, videoDir))
	if err != nil {
		return nil, err
	}
	for _, f := range videos {
		ext := filepath.Ext(f.Name())
		// trim extension
		name := strings.TrimSuffix(f.Name(), ext)
		name = strings.Replace(name, "__", " ", -1)

		v := video{
			Title:  name,
			Source: strings.TrimPrefix(ext, "."),
			Id:     join(videoDir, f.Name()),
		}

		m.Videos = append(m.Videos, v)
	}

	return &m, nil
}

// generateZip writes all files in images to a zip file and returns the file data.
// path is the output path
func (m media) generateZip(outputPath string) error {

	zipOutputPath := join(outputPath, "zip")

	zips := map[string][]string{
		"images.zip": m.Images,
		"logo.zip":   []string{m.Logo, m.Icon},
	}

	for name, fileList := range zips {
		// generate
		zipData, err := zipFiles(outputPath, fileList)
		if err != nil {
			return err
		}
		// safe
		err = writeFile(join(zipOutputPath, name), zipData)
		if err != nil {
			return err
		}
	}

	return nil
}

// used in template to determine if a logo or an icon is present
// TODO can be put in template?
func (m media) HasLogoOrIcon() bool {
	if m.Logo == "" && m.Icon == "" {
		return false
	}
	return true
}
