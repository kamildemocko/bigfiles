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

func GetFiles(folder string, values map[string]File, limit int) error {
	gray := color.New(color.FgHiBlack)
	gray.Print("scanning files...\n\n")

	var smallestKey string
	var smallestSize int64

	err := filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		if len(values) >= limit {
			// remove smallest from map
			delete(values, smallestKey)

			// set next smallest
			smallestKey, smallestSize = findNewSmallest(values)
		}

		name, size := d.Name(), info.Size()

		// add new value
		values[d.Name()] = File{
			Name: name,
			Path: path,
			Size: size,
		}

		// change smallest vars if needed
		if size < smallestSize || smallestKey == "" {
			smallestKey = name
			smallestSize = size
		}

		return err
	})
	if err != nil {
		return err
	}

	return nil
}

func findNewSmallest(values map[string]File) (string, int64) {
	var minSize int64 = 9223372036854775807
	var minSizeKey = ""

	for key, value := range values {
		if value.Size < minSize {
			minSize = value.Size
			minSizeKey = key
		}
	}

	return minSizeKey, minSize
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
