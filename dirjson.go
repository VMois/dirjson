package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fatih/color"
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

func main() {
	dirPath := flag.String("d", ".", "a directory")
	prettyJSON := flag.Bool("p", false, "a pretty json")
	flag.Parse()

	mainDir := NewDirectory(*dirPath)

	files, err := ioutil.ReadDir(*dirPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			mainDir.Dirs = append(mainDir.Dirs, NewDirectory(file.Name()))
		} else {
			mainDir.Files = append(mainDir.Files, File{file.Name(), file.Size()})
		}
	}
	var b []byte
	if *prettyJSON {
		b, err = json.MarshalIndent(mainDir, "", "  ")
	} else {
		b, err = json.Marshal(mainDir)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	greenOutput := color.New(color.FgGreen)
	greenOutput.Println(string(b))
}
