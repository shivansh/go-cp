package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"go-cp/utils"
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
		} else if !srcStat.Mode().IsRegular() {
			// Cannot copy non-regular files apart from
			// directories (e.g. symlinks, devices etc.)
			log.Fatalf("Ignoring non-regular source file %s (%q)",
				   dstStat.Name(), dstStat.Mode().String())
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

// CopyDir recursively copies all the files from the
// directory named 'src' to the directory named 'dst'.
func CopyDir(src, dst string) error {
	src = util.CheckTrailingSlash(src)
	dst = util.CheckTrailingSlash(dst)

	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	var f1, f2 string
	for _, f := range files {
		f1 = src + f.Name()
		f2 = dst + f.Name()
		srcStat, err := os.Stat(f1)
		if err != nil {
			return err
		}
		if srcStat.IsDir() {
			_, err := os.Stat(f2)
			if os.IsNotExist(err) {
				err := os.Mkdir(f2, os.ModePerm)
				if err != nil {
					log.Fatal(err)
				}
			} else if err != nil {
				log.Fatal(err)
			}
			err = CopyDir(f1, f2)
		} else if srcStat.Mode().IsRegular() {
			err = CopyFile(f1, f2)
		} else {
			log.Fatalf("Ignoring non-regular source file %s (%q)",
				   srcStat.Name(), srcStat.Mode().String())
		}
	}
	return nil
}

// CopyFile copies the contents of file named 'src' to the file named 'dst'.
// The destination file will be created if it does not exist. If it already
// exists, then its contents will be replaced by the contents of source file.
func CopyFile(src, dst string) error {
	from, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		return err
	}

	// Flush to disk.
	err = to.Sync()
	if err != nil {
		return err
	}
	return to.Close()  // TODO from.close() ?
}
