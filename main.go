package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// 	"path/filepath"
// 	"strings"
// )

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

func dirTree(out io.Writer, path string, printFiles bool) error {
	return PrintTree(out, path, "", printFiles)
}

func PrintTree(out io.Writer, path, prefix string, printFiles bool) error {
	f, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	file := make([]os.DirEntry, 0, len(f))

	for _, i := range f {
		if i.Name() == ".git" {
			continue
		}
		if i.Name() == "dockerfile" {
			continue
		}
		if i.Name() == "go.mod" {
			continue
		}
		if i.Name() == "readme.md" {
			continue
		}
		if i.Name() == "main.go" {
			continue
		}
		if i.Name() == "main_test.go" {
			continue
		}
		file = append(file, i)
	}

	sort.Slice(file, func(i, j int) bool {
		return file[i].Name() < file[j].Name()
	})

	for index, i := range file {
		isLast := index == len(file)-1

		connect := "├───"
		newprefix := prefix + "│\t"

		if isLast {
			connect = "└───"
			newprefix = prefix + "\t"
		}

		if printFiles {
			if i.IsDir() {
				fmt.Fprintf(out, "%s%s%s\n", prefix, connect, i.Name())
				err := PrintTree(out, filepath.Join(path, i.Name()), newprefix, printFiles)
				if err != nil {
					return err
				}
			} else {
				info, err := i.Info()
				if err != nil {
					return err
				}
				size := info.Size()
				sizeStr := fmt.Sprintf("%db", size)
				if size == 0 {
					sizeStr = "empty"
				}
				fmt.Fprintf(out, "%s%s%s (%s)\n", prefix, connect, i.Name(), sizeStr)
			}
		} else {
			if i.IsDir() {
				fmt.Fprintf(out, "%s%s%s\n", prefix, connect, i.Name())
				err := PrintTree(out, filepath.Join(path, i.Name()), newprefix, printFiles)
				if err != nil {
					return err
				}
			}
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
