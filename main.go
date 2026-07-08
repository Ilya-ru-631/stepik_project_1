package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// 	"path/filepath"
// 	"strings"
// )

import (
	"io"
	"os"
)

func dirTree(r io.Reader, path string, printFiles bool) error {
	txt, err := os.OpenFile(path, os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	defer txt.Close()

	info, err := os.Stat(path)
	if err != nil {
		return nil
	}

	filesize := info.Size()
	buf := make([]byte, filesize)

	if printFiles {
		_, err := txt.WriteString(string(buf))
		if err != nil {
			return nil
		}
	}

	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
