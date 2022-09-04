package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// TODO add depth flag
// TODO add limit on number of files to print
// TODO follow sym links

func tree(out io.Writer, root string, indent string) error {
	fi, err := os.Stat(root)
	if err != nil {
		return err
	}

	fmt.Fprintf(out, "%s\n", fi.Name())

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

		fmt.Fprintf(out, "%s", indent+prefix)

		if err = tree(out, filepath.Join(root, e.Name()), indent+indentSuffix); err != nil {
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

	fmt.Printf("tree %s:\n", dir)
	if err := tree(os.Stdout, dir, ""); err != nil {
		log.Fatal(err)
	}
}
