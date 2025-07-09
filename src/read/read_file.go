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
	WordCount := 0
	LineCount := 0
	var TempWord []rune
	var TempLine []rune
	for RuneCount := 0; RuneCount < len(Contents); RuneCount++ {
		if Contents[RuneCount] == '\n' {
			FileContentsPointer.Words = append(FileContentsPointer.Words,string(TempWord))
			FileContentsPointer.Lines = append(FileContentsPointer.Lines,string(TempLine))
			FileContentsPointer.RuneCount += 1
			fmt.Printf("%s\n",FileContentsPointer.Lines[LineCount])
			LineCount += 1
			WordCount += 1
			TempLine = TempLine[:0]
		}else if Contents[RuneCount] == ' ' {
			FileContentsPointer.Runes = append(FileContentsPointer.Runes,rune(Contents[RuneCount]))
			FileContentsPointer.Words = append(FileContentsPointer.Words,string(TempWord))
			TempLine = append(TempLine,rune(Contents[RuneCount]))
			FileContentsPointer.RuneCount += 1
			WordCount += 1
			TempWord = TempWord[:0]
		}else if Contents[RuneCount] != 0 {
			FileContentsPointer.Runes = append(FileContentsPointer.Runes,rune(Contents[RuneCount]))
			TempWord = append(TempWord,rune(Contents[RuneCount]))
			TempLine = append(TempLine,rune(Contents[RuneCount]))
			FileContentsPointer.RuneCount += 1
			//fmt.Printf("%c ",FileContentsPointer.Runes[RuneCount])
		}
	}
	fmt.Printf("%d",FileContentsPointer.RuneCount)
}
