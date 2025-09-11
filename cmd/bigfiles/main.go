package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"

	"github.com/fatih/color"
)

type File struct {
	Name string
	Path string
	Size int64
}

func getFiles(folder string, values map[string]File) error {
	err := filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		values[d.Name()] = File{
			Name: d.Name(),
			Path: filepath.Join(folder, d.Name()),
			Size: info.Size(),
		}

		return err
	})
	if err != nil {
		return err
	}

	return nil
}

func sortFilesBySize(values map[string]File, asc bool) []string {
	keys := make([]string, 0, len(values))

	for k := range values {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		if asc {
			return values[keys[i]].Size < values[keys[j]].Size
		} else {
			return values[keys[i]].Size > values[keys[j]].Size
		}
	})

	return keys
}

func prettyPrintSize(size int64) string {
	var value float32

	switch {
	case size < 1024:
		value = float32(size)
		return fmt.Sprintf("%.2fB", value)
	case size < 1024*1024:
		value = float32(size) / 1024.0
		return fmt.Sprintf("%.2fKB", value)
	case size < 1024*1024*1024:
		value = float32(size) / (1024.0 * 1024.0)
		return fmt.Sprintf("%.2fMB", value)
	default:
		value = float32(size) / (1024.0 * 1024.0 * 1024.0)
		return fmt.Sprintf("%.2fGB", value)
	}
}

func printResultRow(row File) {
	red := color.New(color.FgHiRed)
	white := color.New(color.FgWhite)
	gray := color.New(color.FgHiBlack)

	red.Printf("[%s] ", prettyPrintSize(row.Size))
	white.Printf("%s", row.Name)
	gray.Printf(" » %s\n", row.Path)
}

func main() {
	folder := `D:\Google Drive\Docs\fonts\`
	values := make(map[string]File)

	err := getFiles(folder, values)
	if err != nil {
		panic(err)
	}

	sorted := sortFilesBySize(values, false)

	limitedAndSorted := sorted[:5]

	for _, sortedKey := range limitedAndSorted {
		row := values[sortedKey]
		printResultRow(row)
	}
}
