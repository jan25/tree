package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// TODO add depth flag
// TODO add limit to number of files to print
// TODO follow sym links

func main() {
	dirs := []string{"."}
	if len(os.Args) > 1 {
		dirs = os.Args[1:]
	}

	// TODO walk all dirs
	dir := dirs[0]
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		seps := strings.Count(strings.TrimPrefix(path, dir), string(filepath.Separator))
		indent := strings.Repeat("  ", seps)
		fmt.Printf("%s%s\n", indent, d.Name())
		return nil
	})

	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
