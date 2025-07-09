package read_file

import "os"
import "fmt"

type FileContents struct {
	RuneCount int
	WordCount int
	LineCount int
	Runes []rune
	Words []string
	Lines []string
}

func ReadFile(FilePath string) []byte{
	Contents := make([]byte, 1024)
	file, err := os.Open("plaintext.txt")
	if err != nil {
		fmt.Println("wrong")
	}
	defer file.Close()
	n, err:= file.Read(Contents)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("wrong2")
	}
	fmt.Println(Contents[:n])
	fmt.Println("\n");
	return Contents
}

func GrabRunes (FileContentsPointer *FileContents,Contents []byte) {
	RuneCount := 0
	WordCount := 0
	LineCount := 0
	var TempWord []string
	var TempLine []string
	for RuneCount := 0; RuneCount < len(Contents); RuneCount++ {
		if Contents [RuneCount] == ' ' {
			FileContentsPointer.Words = append(FileContentsPointer.Words,TempWord)
			WordCount += 1
		}
		if Contents[RuneCount] != 0 {
			FileContentsPointer.Runes = append(FileContentsPointer.Runes,rune(Contents[RuneCount]))
			TempWord = append(TempWord,rune(Contents[RuneCount]))
			TempLine = append(TempLine,rune(Contents[RuneCount]))
			FileContentsPointer.RuneCount += 1
			fmt.Printf("%c ",FileContentsPointer.Runes[RuneCount])
		}
	}
	fmt.Printf("%d",FileContentsPointer.RuneCount)
}
