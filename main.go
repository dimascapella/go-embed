package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version1 string

//go:embed hydra.png
var logo1 []byte

//go:embed files/*.txt
var path1 embed.FS

func main() {
	fmt.Println(version1)

	err := ioutil.WriteFile("logo_new.png", logo1, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dir, _ := path1.ReadDir("files")
	for _, file := range dir {
		if !file.IsDir() {
			fls, _ := path1.ReadFile("files/" + file.Name())
			fmt.Println(string(fls))
		}
	}
}
