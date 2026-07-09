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
	"io/fs"
	"os"
	"path/filepath"
)

func PrintList(path string, printFiles bool) ([]string, error) {
	res := []string{}
	root := path
	if printFiles {
		err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.Name() == ".git" {
				return filepath.SkipDir
			}

			inf, err := d.Info()
			if err != nil {
				return err
			}

			sizee := inf.Size()
		
			if d.IsDir() {
				res = append(res, d.Name())
				//fmt.Println(d.Name())
			} else {
				if sizee == 0 {
					temp_str := fmt.Sprintf("\n%v - %s\n", d.Name(), "empty")
					res = append(res, temp_str)
					//fmt.Printf("\n%v - %s\n", d.Name(), "empty")
				} else {
					temp_str := fmt.Sprintf("\n%v - %d\n", d.Name(), sizee)
					res = append(res, temp_str)
					//fmt.Printf("\n%v - %d\n", d.Name(), sizee)
				}
			}
			return nil
		})
		if err != nil {
			return res, nil
		}
	} else {
		err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.Name() == ".git" {
				return filepath.SkipDir
			}

			if d.IsDir() {
				res = append(res, d.Name(), "\n")
				//fmt.Println(d.Name())
			}

			return nil
		})
		if err != nil {
			return res, nil
		}
	}

	return res, nil
}

func dirTree(w io.Writer, path string, printFiles bool) error {
	res, err := PrintList(path, printFiles)
	if err != nil {
		return err
	}
	for _, file := range res {
		if file.
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
