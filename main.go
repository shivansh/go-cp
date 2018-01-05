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
				err := CopyDir(src, dst)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				err := CopyFile(src, dst + src)
				if err != nil {
					log.Fatal(err)
				}
			}
		} else if !srcStat.IsDir() {
			err := CopyFile(src, dst)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatalf("Omitting directory '%s'", dst)
		}
	} else if os.IsNotExist(err) {
		if srcStat.IsDir() {
			err := os.Mkdir(dst, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
			err = CopyDir(src, dst)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := CopyFile(src, dst)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		log.Fatal(err)
	}
}

// CopyDir copies all the files from directory 'src' to directory 'dst'.
func CopyDir(src, dst string) error {
	src = CheckTrailingSlash(src)
	dst = CheckTrailingSlash(dst)

	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, f := range files {
		err := CopyFile(src + f.Name(), dst + f.Name())
		if err != nil {
			return err
		}
	}
	return nil
}

// CopyFile copies the contents of file 'src' to file 'dst'.
func CopyFile(src, dst string) error {
	from, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		return err
	}
	return to.Close()  // TODO from.close() ?
}

// CheckTrailingSlash checks the presence of trailing slash in
// case the argument is a directory, and appends on if absent.
func CheckTrailingSlash(dir string) string {
	if dir[len(dir) - 1] != '/' {
		dir += "/"
	}
	return dir
}
