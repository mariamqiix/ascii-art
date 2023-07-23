package main

import (
	"ascii"
	"fmt"
	"os"
	"strings"
)

func main() {
	Flag := ascii.Validation()
	if Flag {
		WordsInArr := strings.Split(os.Args[1], "\\n")
		for l := 0; l < len(WordsInArr); l++ {
			var Words [][]string
			Text1 := WordsInArr[l]
			if string(Text1) == "" {
				fmt.Println("")
			} else {
				for j := 0; j < len(Text1); j++ {
					Words = append(Words, ascii.ReadLetter(Text1[j]))
				}
				for w := 0; w < 8; w++ {
					for n := 0; n < len(Words); n++ {
						fmt.Print(Words[n][w])
					}
					if w+1 != 8 {
						fmt.Println()
					}
				}
				fmt.Println()
			}
		}
	}
}
