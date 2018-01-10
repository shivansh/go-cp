package test

import (
	"log"
	"os"
	"testing"
)

func setup() {
	err := os.MkdirAll("srcdir/a", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	_, err = os.Create("srcdir/b")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir("dstdir", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	_, err = os.Create("dstdir/b")
	if err != nil {
		log.Fatal(err)
	}
}

func teardown() {
	err := os.RemoveAll("srcdir")
	if err != nil {
		log.Fatal(err)
	}
	err = os.RemoveAll("dstdir")
	if err != nil {
		log.Fatal(err)
	}
}

func TestCompare(t *testing.T) {
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{{
		name: "Directory",
		args: args{"srcdir", "dstdir"},
		want: false,
	},
	{
		name: "File",
		args: args{"srcdir/b", "dstdir/b"},
		want: true,
	}}
	for _, tt := range tests {
		if got := Compare(tt.args.src, tt.args.dst); got != tt.want {
			t.Errorf("%q. Compare() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestMain(m *testing.M) {
	setup()
	retcode := m.Run()
	teardown()
	os.Exit(retcode)
}
