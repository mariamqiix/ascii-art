package main 

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	Text1 := os.Args[1]
	if len(os.Args) != 2 {
		fmt.Println("Error")
		os.Exit(0)
	}
	ReadFile, _ := os.Open("standard.txt")
	FileScanner := bufio.NewScanner(ReadFile)
	FileScanner.Split(bufio.ScanLines)

    var Words [][]string

	for j := 0 ; j< len(Text1); j++ {
		var Letter []string
		stop := 1
		character := int(Text1[j])
		i := 0
		a := character - 32
		a = a*9 +2
		for FileScanner.Scan() {
			i++
			if i >= a {
				fmt.Println(character)

				stop++
				Letter = append(Letter, FileScanner.Text())
				if stop > 8 {
					break
				}
			}
		}
		Words = append(Words, Letter)
	}
	for h:=0; h<len(Words); h++ {
		for _, line := range Words[h] {
			fmt.Println(line)
		}
	}

    ReadFile.Close()
}