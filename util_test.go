package presskit

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestJoin(t *testing.T) {
	parts := []string{"this", "is", "my", "path"}
	j := join(parts...)
	fj := filepath.Join(parts...)
	if j != fj {
		t.Error("Join doesn't work properly")
	}
}

func TestZipFiles(t *testing.T) {
	basepath := "./test/zipfiles"
	testArchive := filepath.Join(basepath, "archive.zip")

	// return an error on not existing file
	_, err := zipFiles(basepath, []string{"notexisting"})
	if err == nil {
		t.Error("zipFiles should return an error on a nonexistant file")
	}

	// create and save zip
	zipData, err := zipFiles(basepath, []string{"file1.txt", "file2.txt"})
	if err != nil {
		t.Error("error in zipFiles:" + err.Error())
	}

	err = ioutil.WriteFile(testArchive, zipData, 0644)
	if err != nil {
		t.Error("error writing test archive: " + err.Error())
	}

	// read the file
	r, err := zip.OpenReader(testArchive)
	if err != nil {
		t.Error("error opening test archive: " + err.Error())
	}
	defer r.Close()

	for _, f := range r.File {
		// get content of zipped file
		var zipContent bytes.Buffer
		fileZipped, err := f.Open()
		if err != nil {
			t.Error(err)
		}
		_, err = io.Copy(&zipContent, fileZipped)
		if err != nil {
			t.Error(err)
		}
		fileZipped.Close()

		// get content of original file
		originalContent, err := ioutil.ReadFile(filepath.Join(basepath, f.Name))
		if err != nil {
			t.Error(err)
		}

		// compare
		if !reflect.DeepEqual(zipContent.Bytes(), originalContent) {
			t.Error("zipped content doesn't equal original content")
		}
	}

	// cleanup
	if err := os.Remove(testArchive); err != nil {
		t.Error("error removing test archive: " + err.Error())
	}
}
