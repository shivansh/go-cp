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
			// TODO Validate if recursive copying is the,
			// intended behavior, for e.g. via '-r' option.
			if dst[len(dst) - 1] != '/' {
				dst += "/"
			}

			if srcStat.IsDir() {
				files, err := ioutil.ReadDir(src)
				if err != nil {
					log.Fatal(err)
				}

				for _, f := range files {
					Copy(src + "/" + f.Name(), dst + f.Name())
				}
			} else {
				Copy(src, dst + src)
			}
		} else if !srcStat.IsDir() {
			Copy(src, dst)
		} else {
			log.Fatalf("Omitting directory '%s'", dst)
		}
	} else if os.IsNotExist(err) {
		if srcStat.IsDir() {
			err := os.Mkdir(dst, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}

			files, err := ioutil.ReadDir(src)
			if err != nil {
				log.Fatal(err)
			}

			for _, f := range files {
				Copy(src + "/" + f.Name(), dst + "/" + f.Name())
			}
		} else {
			Copy(src, dst)
		}
	} else {
		log.Fatal(err)
	}
}

func Copy(src, dst string) {
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
