package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/fatih/color"
)

// File struct that represents file
type File struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

// Directory struct that represents dir
type Directory struct {
	Name  string      `json:"name"`
	Dirs  []Directory `json:"dirs"`
	Files []File      `json:"files"`
}

// NewDirectory creates new Directory struct
func NewDirectory(path string) Directory {
	newDir := Directory{}
	newDir.Name = path
	newDir.Dirs = []Directory{}
	newDir.Files = []File{}
	return newDir
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	dirPath := flag.String("d", ".", "a directory")
	prettyJSON := flag.Bool("p", false, "a pretty JSON output")
	outputFlag := flag.String("o", "", "save output to file (filename)")
	flag.Parse()

	currentRunDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	check(err)

	mainDir := NewDirectory(*dirPath)

	files, err := ioutil.ReadDir(*dirPath)
	check(err)

	for _, file := range files {
		if file.IsDir() {
			mainDir.Dirs = append(mainDir.Dirs, NewDirectory(file.Name()))
		} else {
			mainDir.Files = append(mainDir.Files, File{file.Name(), file.Size()})
		}
	}

	// convert to json
	var b []byte
	if *prettyJSON {
		b, err = json.MarshalIndent(mainDir, "", "  ")
	} else {
		b, err = json.Marshal(mainDir)
	}
	check(err)

	// save output to file or print
	if *outputFlag != "" {
		outputDir := path.Join(currentRunDir, *outputFlag)
		f, err := os.Create(outputDir)
		check(err)
		defer f.Close()
		_, err = f.Write(b)
		check(err)
	} else {
		greenOutput := color.New(color.FgHiGreen)
		greenOutput.Println(string(b))
	}
}
