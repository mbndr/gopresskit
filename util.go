package presskit

import (
	"archive/zip"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func html(format string, a ...interface{}) template.HTML {
	return template.HTML(fmt.Sprintf(format, a...))
}

// copyDir copies all files in a directory (not recursively)
func copyDir(source, destination string) error {
	// get source properties
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create destination directory
	err = os.MkdirAll(destination, sourceinfo.Mode())
	if err != nil {
		return err
	}

	// read source directory
	files, err := ioutil.ReadDir(source)
	if err != nil {
		return err
	}

	for _, f := range files {
		// open source file
		sf, err := os.Open(filepath.Join(source, f.Name()))
		if err != nil {
			return err
		}
		// create destination file
		df, err := os.Create(filepath.Join(destination, f.Name()))
		if err != nil {
			return err
		}
		// copy
		_, err = io.Copy(df, sf)
		if err != nil {
			return err
		}
	}

	return nil
}

// zipFiles zips a list of files and returns the zips data
func zipFiles(basepath string, files []string) ([]byte, error) {
	buf := new(bytes.Buffer)

	z := zip.NewWriter(buf)

	for _, p := range files {
		fullPath := join(basepath, p)
		// add to zip
		f, err := z.Create(path.Base(p))
		if err != nil {
			return nil, err
		}
		// read file data
		fileData, err := ioutil.ReadFile(fullPath)
		if err != nil {
			return nil, err
		}
		// write file data to zip
		_, err = f.Write(fileData)
		if err != nil {
			return nil, err
		}
	}

	err := z.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// writeFile is a shortcut because of file permissions
func writeFile(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0644)
}

// join is a shortcut for filepath.Join
func join(p ...string) string {
	return filepath.Join(p...)
}
