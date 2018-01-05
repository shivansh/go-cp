package main

import (
	"log"
	"os"
	"testing"
	"go-cp/test"
)

func setup() {
	err := os.Mkdir("tests/dstdir", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func teardown() {
	err := os.RemoveAll("tests/dstfile")
	if err != nil {
		log.Fatal(err)
	}
	err = os.RemoveAll("tests/dstdir")
	if err != nil {
		log.Fatal(err)
	}
}

func TestCopyDir(t *testing.T) {
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{{
		name: "Directory",
		args: args{"tests/srcdir", "tests/dstdir"},
		wantErr: false,
	},
	}
	for _, tt := range tests {
		if err := CopyDir(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
			t.Errorf("%q. CopyDir() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestCopyFile(t *testing.T) {
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{{
		name: "TextFile",
		args: args{"tests/srcfile", "tests/dstfile"},
		wantErr: false,
	},
	}
	for _, tt := range tests {
		if err := CopyFile(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
			t.Errorf("%q. CopyFile() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		} else if !test.Compare(tt.args.src, tt.args.dst) {
			t.Errorf("%q. CopyFile() error = %v, wantErr %v", tt.name, err, tt.wantErr)

		}
	}
}

func TestMain(m *testing.M) {
	setup()
	retcode := m.Run()
	teardown()
	os.Exit(retcode)
}
