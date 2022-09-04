package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// TODO add depth flag
// TODO add limit on number of files to print
// TODO follow sym links

func tree2(dir string) error {
	return filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == dir {
			fmt.Printf("%s\n", dir)
			return nil
		}

		seps := strings.Count(strings.TrimPrefix(path, dir), string(filepath.Separator))
		indent := strings.Repeat("   ", seps-1) + "├── "
		if seps > 1 {
			indent = "│" + indent
		}
		fmt.Printf("%s%s\n", indent, d.Name())
		return nil
	})
}

func tree(root string, indent string) error {
	fi, err := os.Stat(root)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", fi.Name())

	if !fi.IsDir() {
		return nil
	}

	entries, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	for i, e := range entries {
		prefix := "├── "
		indentSuffix := "    "
		if i == len(entries)-1 {
			prefix = "└── "
		} else {
			indentSuffix = "│   "
		}

		fmt.Printf("%s", indent+prefix)

		if err = tree(filepath.Join(root, e.Name()), indent+indentSuffix); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	dirs := []string{"."}
	if len(os.Args) > 1 {
		dirs = os.Args[1:]
	}

	// TODO walk all dirs
	dir := dirs[0]
	if err := tree(dir, ""); err != nil {
		log.Fatal(err)
	}
}
