package presskit

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"code.cloudfoundry.org/bytefmt"
)

const (
	// directory of images
	imageDir = "images"
	// directory of videos
	videoDir = "videos"
)

// media holds all media information of a document
type media struct {
	Header   string
	Icon     string
	Logo     string
	Images   []string // images for the images section
	Videos   []video
	zipSizes map[string]string
}

// newMedia reads all images and videos
func newMedia(path string, h videoHolder) (*media, error) {
	var m media

	// images
	images, _ := ioutil.ReadDir(join(path, imageDir))
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
	m.Videos = h.GetVideos()

	// add local videos (err if no folder)
	videos, _ := ioutil.ReadDir(join(path, videoDir))
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

	m.zipSizes = make(map[string]string)

	return &m, nil
}

// generateZip writes all files in images to a zip file and returns the file data
func (m media) generateZip(outputPath string) error {

	zipOutputPath := join(outputPath, "zip")

	zips := map[string][]string{
		"images.zip": m.Images,
		"logo.zip":   {m.Logo, m.Icon},
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

		// TODO own format func
		m.zipSizes[name] = bytefmt.ByteSize(uint64(len(zipData)))
	}

	return nil
}

// ZipSize is called from template to get the size string representation of a zip file
func (m media) ZipSize(name string) string {
	return m.zipSizes[name]
}

// HasLogoOrIcon is used in template to determine if a logo or an icon is present
func (m media) HasLogoOrIcon() bool {
	if m.Logo == "" && m.Icon == "" {
		return false
	}
	return true
}

// videoHolder is used to get the videos from both game and company in newMedia
type videoHolder interface {
	GetVideos() []video
}

// GetVideos returns all videos of the data file
func (c company) GetVideos() []video {
	return c.Videos
}

// GetVideos returns all videos of the data file
func (g game) GetVideos() []video {
	return g.Videos
}
