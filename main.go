package main

import (
	//"fmt"
	"github.com/ydnawizard/AES/src/read"
)


func main() {
	//var FilePath string
	FileContentsPointer := new(read_file.FileContents)
	Contents := read_file.ReadFile("plaintext.txt")
	read_file.GrabRunes(FileContentsPointer,Contents)
}

