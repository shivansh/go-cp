package test

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"go-cp/utils"
)

// Compare checks whether src and dst have same content.
func Compare(src, dst string) bool {
	srcStat, err := os.Stat(src)
	if err != nil {
		log.Fatal(err)
	}

	dstStat, err := os.Stat(dst)
	if err != nil {
		log.Fatal(err)
	}

	if srcStat.IsDir() && dstStat.IsDir() {
		return CompareDir(src, dst)
	} else if !srcStat.IsDir() && !dstStat.IsDir() {
		return CompareFile(src, dst)
	} else {
		return false
	}
}

// CompareDir checks if the directories passed as arguments have same content.
func CompareDir(dir1, dir2 string) bool {
	var f1 []string
	var f2 []string

	dir1files, err := ioutil.ReadDir(dir1)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range dir1files {
		f1 = append(f1, f.Name())
	}

	dir2files, err := ioutil.ReadDir(dir2)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range dir2files {
		f2 = append(f2, f.Name())
	}

	if len(f1) != len(f2) {
		return false
	}

	dir1 = util.CheckTrailingSlash(dir1)
	dir2 = util.CheckTrailingSlash(dir2)
	for i := 0; i < len(f1); i++ {
		if !Compare(dir1 + f1[i], dir2 + f2[i]) {
			return false
		}
	}
	return true
}

// Compare checks if the files passed as arguments have same content.
func CompareFile(file1, file2 string) bool {
	// Minimum granularity (in bytes) at which the files are compared.
	chunkSize := 32 * 1024

	// Compare file sizes.
	f1stat, err := os.Lstat(file1)
	if err != nil {
		log.Fatal(err)
	}

	f2stat, err := os.Lstat(file2)
	if err != nil {
		log.Fatal(err)
	}

	if f1stat.Size() != f2stat.Size() {
		return false
	}

	f1, err := os.Open(file1)
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	// Compare file contents chunk-by-chunk.
	for {
		chunk1 := make([]byte, chunkSize)
		_, err1 := f1.Read(chunk1)

		chunk2 := make([]byte, chunkSize)
		_, err2 := f2.Read(chunk2)

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true
			} else if err1 == io.EOF || err2 == io.EOF {
				return false
			} else {
				log.Fatal(err1, err2)
			}
		}

		if !bytes.Equal(chunk1, chunk2) {
			return false
		}
	}
}
