package ascii

import (
	"bufio"
	"os"
	// "fmt"
)

func ReadLetter(Text1 byte) []string {
	//buffio object, to open and read the file
	ReadFile, _ := os.Open("standard.txt")
	FileScanner := bufio.NewScanner(ReadFile)
	var Letter []string
	stop := 1
	i := 0
	a := (int(Text1)-32)*9 + 2
	for FileScanner.Scan() {
		i++
		if i >= a {
			stop++
			Letter = append(Letter, FileScanner.Text())
			if stop > 8 {
				break
			}
		}
	}
	ReadFile.Close()
	return Letter
}
