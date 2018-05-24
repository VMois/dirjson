package lib

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

// File struct that represents file
type File struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

// Directory struct that represents dir
type Directory struct {
	Path  string      `json:"path"`
	Dirs  []Directory `json:"dirs"`
	Files []File      `json:"files"`
}

// NewDirectory creates new Directory struct
func NewDirectory(path string) Directory {
	newDir := Directory{}
	newDir.Path = path
	newDir.Dirs = []Directory{}
	newDir.Files = []File{}
	return newDir
}

// DirsExplorer scan directories
func DirsExplorer(rootDir *Directory, recursiveScan bool) {
	files, err := ioutil.ReadDir(rootDir.Path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			newDir := NewDirectory(filepath.Join(rootDir.Path, file.Name()))
			if recursiveScan {
				DirsExplorer(&newDir, true)
			}
			rootDir.Dirs = append(rootDir.Dirs, newDir)
		} else {
			rootDir.Files = append(rootDir.Files, File{file.Name(), file.Size()})
		}
	}
}
