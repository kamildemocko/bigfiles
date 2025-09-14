package printer

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/kamildemocko/bigfiles/internal/operations"
)

type Printer struct {
	Dir string
}

var (
	red   = color.New(color.FgHiRed)
	white = color.New(color.FgWhite)
	gray  = color.New(color.FgHiBlack)
)

func NewPrinter(rootDir string) Printer {
	return Printer{
		Dir: rootDir,
	}
}

func (p *Printer) PrintResultRow(row operations.File) {
	red.Printf("[ %s ] ", prettyPrintSize(row.Size))
	white.Printf("%s", row.Name)
	path := strings.Replace(row.Path, p.Dir, "", 1)
	path = filepath.Dir(path)
	gray.Printf(" » %s\n", path)
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
