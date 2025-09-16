package operations

import (
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

func GetFiles(folder string, values map[string]File) error {
	gray := color.New(color.FgHiBlack)
	gray.Print("scanning files...\n\n")

	err := filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		info, err := d.Info()
		if err != nil {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		values[d.Name()] = File{
			Name: d.Name(),
			Path: path,
			Size: info.Size(),
		}

		return err
	})
	if err != nil {
		return err
	}

	return nil
}

func SortFilesBySize(values map[string]File, asc bool) []string {
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
