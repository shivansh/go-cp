package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Usage: go-cp SOURCE DEST")
		os.Exit(1)
	}

	src := os.Args[1]
	dst := os.Args[2]

	srcStat, err := os.Stat(src)
	if err != nil {
		log.Fatal(err)
	}

	dstStat, err := os.Stat(dst)
	if err == nil {
		if dstStat.IsDir() {
			// TODO Validate if recursive copying is the
			// intended behavior, for e.g. via '-r' option.
			if srcStat.IsDir() {
				CopyDir(src, dst)
			} else {
				CopyFile(src, dst + src)
			}
		} else if !srcStat.IsDir() {
			CopyFile(src, dst)
		} else {
			log.Fatalf("Omitting directory '%s'", dst)
		}
	} else if os.IsNotExist(err) {
		if srcStat.IsDir() {
			err := os.Mkdir(dst, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
			CopyDir(src, dst)
		} else {
			CopyFile(src, dst)
		}
	} else {
		log.Fatal(err)
	}
}

// Copies all the files from directory 'src' to directory 'dst'.
func CopyDir(src, dst string) {
	src = CheckTrailingSlash(src)
	dst = CheckTrailingSlash(dst)

	files, err := ioutil.ReadDir(src)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		CopyFile(src + f.Name(), dst + f.Name())
	}
}

// Copies the contents of file 'src' to file 'dst'.
func CopyFile(src, dst string) {
	from, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}
}

// In case the argument is a directory, checks the
// presence of trailing slash and appends one if absent.
func CheckTrailingSlash(dir string) string {
	if dir[len(dir) - 1] != '/' {
		dir += "/"
	}
	return dir
}
