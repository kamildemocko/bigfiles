package main

import (
	"github.com/kamildemocko/bigfiles/internal/operations"
	"github.com/kamildemocko/bigfiles/internal/printer"
)

func main() {
	limit := 5
	folder := `D:\Google Drive\Docs\fonts\`
	values := make(map[string]operations.File)

	err := operations.GetFiles(folder, values)
	if err != nil {
		panic(err)
	}

	sorted := operations.SortFilesBySize(values, false)

	limitedAndSorted := sorted[:limit]

	p := printer.NewPrinter(folder)
	for _, sortedKey := range limitedAndSorted {
		row := values[sortedKey]
		p.PrintResultRow(row)
	}
}
