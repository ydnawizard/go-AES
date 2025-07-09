package main

import (
	//"fmt"
	"github.com/ydnawizard/AES/src/read"
)


func main() {
	//var FilePath string
	FileContentsPointer := new(read_file.FileContents)
	read_file.ReadFileContents(FileContentsPointer,"plaintext.txt")
}

