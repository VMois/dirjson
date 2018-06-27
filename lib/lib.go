package lib

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"
)

// File struct that represents file
type File struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

// Directory struct that represents dir
type Directory struct {
	Path  string       `json:"path"`
	Dirs  *[]Directory `json:"dirs"`
	Files *[]File      `json:"files"`
}

// NewDirectory creates new Directory struct
func NewDirectory(path string) Directory {
	newDir := Directory{}
	newDir.Path = path
	newDir.Dirs = &[]Directory{}
	newDir.Files = &[]File{}
	return newDir
}

// DirsExplorer scan directories
func DirsExplorer(rootDir *Directory, recursiveScan bool) {
	var wg sync.WaitGroup
	wg.Add(1)
	go dirsRunner(rootDir, recursiveScan, &wg)
	wg.Wait()
}

func dirsRunner(rootDir *Directory, recursiveScan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	files, err := ioutil.ReadDir(rootDir.Path)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	for _, file := range files {
		if file.IsDir() {
			newDir := NewDirectory(filepath.Join(rootDir.Path, file.Name()))
			*rootDir.Dirs = append(*rootDir.Dirs, newDir)
			if recursiveScan {
				wg.Add(1)
				go dirsRunner(&newDir, recursiveScan, wg)
			}
		} else {
			*rootDir.Files = append(*rootDir.Files, File{file.Name(), file.Size()})
		}
	}
}
