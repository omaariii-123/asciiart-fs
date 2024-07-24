package main

import (
	"fmt"
	"os"

	ascii "ascii/function"
)

func main() {
	if len(os.Args) == 3 {
		if os.Args[1] != "" && os.Args[2] != "" {
			ascii.Protect(os.Args[1], &os.Args[2], 2)
			amap := ascii.ParseFile(os.Args[2])
			ascii.PrintAscii(os.Args[1], amap)
		}
	} else if len(os.Args) == 2 {
		ascii.Protect(os.Args[1], nil, 1)
		amap := ascii.ParseFile("standard")
		ascii.PrintAscii(os.Args[1], amap)
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
	}
}
