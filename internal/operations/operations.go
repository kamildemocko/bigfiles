package operations

import (
	"io/fs"
	"path/filepath"
	"sort"
)

type File struct {
	Name string
	Path string
	Size int64
}

func GetFiles(folder string, values map[string]File) error {
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
