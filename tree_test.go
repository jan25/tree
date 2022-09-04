package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

var paths = []string{
	"/a/foo",
	"/b/bar",
	"/meh",
	"/c/d/baz",
}

const expected = `%s
├── a
│   └── foo
├── b
│   └── bar
├── c
│   └── d
│       └── baz
└── meh
`

func getExpected(root string) string {
	b := filepath.Base(root)
	return fmt.Sprintf(expected, b)
}

func prepTempFs(root string) error {
	for _, p := range paths {
		if err := os.MkdirAll(root+p, 0777); err != nil {
			return err
		}
	}
	return nil
}

func TestTree(t *testing.T) {
	root := t.TempDir()

	if err := prepTempFs(root); err != nil {
		t.Errorf("error: %v", err)
	}

	var buf bytes.Buffer
	if err := tree(&buf, root, ""); err != nil {
		t.Errorf("error: %v", err)
	}

	exp := getExpected(root)
	if buf.String() != exp {
		t.Errorf("want:\n%s\ngot:\n%s\n", exp, buf.String())
	} else {
		t.Logf("expected output: %s\n", exp)
	}
}
