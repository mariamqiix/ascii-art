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
		newLine := ascii.OnlyContains(os.Args[1], "\\n")
		WordsInArr := strings.Split(os.Args[1], "\\n")
		fileName := "standard"
		if len(os.Args) == 3 { 
			fileName = strings.ToLower(os.Args[2])
		}
		if newLine == true && len(WordsInArr) != 1 {
			for i := 0 ; i < len(WordsInArr)-1; i++ {
				fmt.Println()
			}
			return
		}
		
		for l := 0; l < len(WordsInArr); l++ {
			var Words [][]string
			Text1 := WordsInArr[l]
			Text1 = strings.ReplaceAll(Text1, "\\t", "   ")
			if string(Text1) == "" {
				fmt.Println("")
			} else {
				for j := 0; j < len(Text1); j++ {
					Words = append(Words, ascii.ReadLetter(Text1[j], fileName ))
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
