package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kamildemocko/bigfiles/internal/operations"
	"github.com/kamildemocko/bigfiles/internal/printer"
)

var (
	rootPath string
	limit    int
)

func init() {
	flag.StringVar(&rootPath, "d", "", "set root directory, if ommited CWD will be used")
	flag.IntVar(&limit, "l", 5, "set max shown files")
	flag.Parse()
}

func parseInputDir() (string, error) {
	var folder string
	var err error

	if rootPath == "" {
		folder, err = os.Getwd()
		if err != nil {
			return "", fmt.Errorf("cannot get current directory")
		}
	} else {
		folder = rootPath
	}
	_, err = os.Stat(folder)
	if err != nil {
		return "", fmt.Errorf("directory not found")
	}

	return folder, nil
}

func main() {
	folder, err := parseInputDir()
	if err != nil {
		fmt.Println(err)
		return
	}
	allFiles := make(map[string]operations.File)

	err = operations.GetFiles(folder, allFiles)
	if err != nil {
		panic(err)
	}

	sorted := operations.SortFilesBySize(allFiles, false)

	limitedAndSorted := sorted[:limit]

	p := printer.NewPrinter(folder)
	for _, sortedKey := range limitedAndSorted {
		row := allFiles[sortedKey]
		p.PrintResultRow(row)
	}
}
