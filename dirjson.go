package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

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

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func dirsExplorer(rootDir *Directory) {
	files, err := ioutil.ReadDir(rootDir.Path)
	check(err)

	for _, file := range files {
		if file.IsDir() {
			newDir := NewDirectory(filepath.Join(rootDir.Path, file.Name()))
			dirsExplorer(&newDir)
			rootDir.Dirs = append(rootDir.Dirs, newDir)
		} else {
			rootDir.Files = append(rootDir.Files, File{file.Name(), file.Size()})
		}
	}
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	dirPath := flag.String("d", ".", "a directory")
	prettyJSON := flag.Bool("p", false, "a pretty JSON output")
	outputFlag := flag.String("o", "", "save output to file (filename)")
	flag.Parse()

	currentRunDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	check(err)

	PrintMemUsage()
	mainDir := NewDirectory(*dirPath)
	dirsExplorer(&mainDir)
	PrintMemUsage()

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
