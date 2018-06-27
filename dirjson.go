package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"path"
	"path/filepath"

	lib "github.com/VMois/dirjson/lib"
	"github.com/fatih/color"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	dirPath := flag.String("d", ".", "a directory")
	prettyJSON := flag.Bool("p", false, "a pretty JSON output")
	outputFlag := flag.String("o", "", "save output to file (filename)")
	recursiveFlag := flag.Bool("r", false, "set recursive directory scan")
	flag.Parse()

	currentRunDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	check(err)

	rootDir := lib.NewDirectory(*dirPath)
	lib.DirsExplorer(&rootDir, *recursiveFlag)

	// convert to json
	var b []byte
	if *prettyJSON {
		b, err = json.MarshalIndent(&rootDir, "", "  ")
	} else {
		b, err = json.Marshal(&rootDir)
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
