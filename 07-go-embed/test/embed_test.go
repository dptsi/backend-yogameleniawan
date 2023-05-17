package test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println("Version:", version)
}

//go:embed logo.png
var logo []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)

	if err != nil {
		t.Error(err)
	}

}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestFS(t *testing.T) {
	a, err := files.ReadFile("files/a.txt")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(a))

	b, err := files.ReadFile("files/b.txt")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(b))

	c, err := files.ReadFile("files/c.txt")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(c))
}

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(content))
		}
	}
}
