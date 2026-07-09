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
	"strings"
)

func depthOf(path string) ([]int, error) {
	res := []int{}
	root := path
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		var depth int
		if rel == "." {
			depth = 0
		} else {
			depth = strings.Count(rel, string(filepath.Separator)) + 1
		}
		res = append(res, depth)
		//fmt.Printf("depth=%d\t%s\n", depth, path)
		return nil
	})
	if err != nil {
		return res, err
	}
	return res, err
}

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

func IsDirectoria(path string) (bool, error) {
	root := path
	flag := true
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			flag = true
		} else {
			flag = false
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return flag, nil
}

func Dep(path, file string, i []int) int {
	ret := 0
	res, _ := depthOf(path)
	rr, _ := PrintList(path, true)
	for i, _ := range rr {
		ret = res[i]
	}
	return ret
}

func dirTree(w io.Writer, path string, printFiles bool) error {
	res, err := PrintList(path, printFiles)
	if err != nil {
		return err
	}

	res2, err := depthOf(path)
	if err != nil {
		return err
	}

	for _, file := range res {
		tr, _ := IsDirectoria(file)
		if tr {
			depDir := Dep(path, file, res2)
			for i := 0; i < depDir; i++ {
				fmt.Println("|")
				if i == depDir-1 {
					fmt.Println("├───")
				}
			}
		} else {
			depFile := Dep(path, file, res2)
			if depFile == 2 {
				fmt.Println("	├───")
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
