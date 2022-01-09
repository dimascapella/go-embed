package go_embed

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
	fmt.Println(version)
}

//go:embed hydra.png
var logo []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/fileA.txt
//go:embed files/fileB.txt
//go:embed files/fileC.txt

var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/fileA.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/fileB.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/fileC.txt")
	fmt.Println(string(c))
}

//go:embed files/*.txt

var path embed.FS

func TestPathMather(t *testing.T) {
	dir, _ := path.ReadDir("files")
	for _, file := range dir {
		if !file.IsDir() {
			fls, _ := path.ReadFile("files/" + file.Name())
			fmt.Println(string(fls))
		}
	}
}
