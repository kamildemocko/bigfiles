package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/kamildemocko/bigfiles/internal/operations"
	"github.com/kamildemocko/bigfiles/internal/printer"
)

var (
	rootPath string
	limit    int
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [DIRECTORY]\n\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(os.Stderr, "If DIRECTORY is not specified, the current directory will be used.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.IntVar(&limit, "l", 5, "set max shown files")
	flag.Parse()

	if len(flag.Args()) > 0 {
		argPath := flag.Args()
		rootPath = strings.Join(argPath, " ")
	}
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
	info, err := os.Stat(folder)
	if err != nil {
		return "", fmt.Errorf("invalid directory")
	}
	if !info.IsDir() {
		gray := color.New(color.FgHiBlack)
		gray.Print("this is a file, using parent directory instead\n\n")
		folder = filepath.Dir(folder)
	}

	return folder, nil
}

func main() {
	folder, err := parseInputDir()
	if err != nil {
		fmt.Println(err)
		return
	}
	allFiles := make(map[string]operations.File, limit+1)

	err = operations.GetFiles(folder, allFiles, limit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	sorted := operations.SortFilesBySize(allFiles, false)
	if len(sorted) == 0 {
		fmt.Println("no files in folder")
		os.Exit(0)
	}
	if limit > len(sorted) {
		limit = len(sorted)
	}
	limitedAndSorted := sorted[:limit]

	p := printer.NewPrinter(folder)
	for _, sortedKey := range limitedAndSorted {
		row := allFiles[sortedKey]
		p.PrintResultRow(row)
	}
}
