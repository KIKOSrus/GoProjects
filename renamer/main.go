package main

// СЕМКА ПОШЕЛ НАХУЙ ЗАЕБАЛ

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var re = regexp.MustCompile(`^(.*)(\d+)\.(\w*)$`)

type file struct {
	Newpath string
	Oldpath string
}

func Change(files []file) error {
	for _, f := range files {
		err := os.Rename(f.Oldpath, f.Newpath)
		if err != nil {
			return err
		}
	}
	return nil
}

func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func match(s string) (string, error) {
	total := 4
	groups := re.FindStringSubmatch(s)
	if len(groups) == 0 {
		return "", nil
	} else {
		return fmt.Sprintf("%s - (%s out of %d).%s", groups[1], groups[2], total, groups[3]), nil
	}
}

func visit(path string, d os.DirEntry, err error) error {
	var toChange []file

	name, err := match(path)

	Must(err)

	if name != "" {
		newfile := file{name, path}
		toChange = append(toChange, newfile)
	}
	err = Change(toChange)

	Must(err)
	return nil
}

func main() {

	// directory where the changes should be made
	root := "."
	var dir string
	_, err := fmt.Scanf("%s\n", &dir)
	Must(err)
	root += string(filepath.Separator) + dir

	// opening a directory
	err = filepath.WalkDir(root, visit)

	Must(err)
}
