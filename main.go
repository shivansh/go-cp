package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"go-cp/test"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Usage: go-cp SOURCE DEST")
		os.Exit(1)
	}

	src := os.Args[1]
	dst := os.Args[2]

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

	// TODO Check if validation is done as a part of "io.Copy".
	if !test.Compare(src, dst) {
		fmt.Println("An error occurred while copying files.")
	} else {
		fmt.Println("File transfer complete.")
	}
}
