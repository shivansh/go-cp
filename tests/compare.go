package test

import (
	"bytes"
	"io"
	"log"
	"os"
)

// Compare checks if the files passed as arguments have same contents.
func Compare(file1, file2 string) bool {
	// Minimum granularity (in bytes) at which the files are compared.
	chunkSize := 1024

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
