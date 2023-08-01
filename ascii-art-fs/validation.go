package ascii

import (
	"fmt"
	"os"
	"strings"
)

func Validation() bool {
	
	//To check the number of arguments
	if len(os.Args) > 3 || len(os.Args) < 2 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		os.Exit(0)
	}
	if len(os.Args) == 3 {
		FontType := strings.ToLower(os.Args[2])
		if FontType != "standard" && FontType != "shadow" && FontType != "thinkertoy" {
			fmt.Println("Usage: go run . [STRING] [BANNER]")
			fmt.Println("EX: go run . something standard")
			os.Exit(0)
		}
	}
	//To validate the character of the strings in ascii range only
	for g := 0; g < len(os.Args[1]); g++ {
		if os.Args[1][g] > 126 || os.Args[1][g] < 32 {
			fmt.Println("ERROR: ascii letters only")
			os.Exit(0)
		}
	}
	if len(os.Args[1]) == 0 {
		return false
	} else if os.Args[1] == "\\n" {
		fmt.Println()
		return false
	}
	return true
}
