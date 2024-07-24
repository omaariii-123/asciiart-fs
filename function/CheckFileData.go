package ascii

import (
	"fmt"
	"os"
	"strings"
)
/* checks whether the content of the files are modifed or not*/
func CheckFileData(s string) {
	SplitedString := strings.Split(s, "\n\n")
	LenSS := len(SplitedString)
	if LenSS != 95 {
		fmt.Println("error : one of the symbols or more are missing from the banner file !")
		os.Exit(1)
	}
	for i := 0; i < LenSS; i++ {
		if len(strings.Split(SplitedString[i], "\n")) != 8 {
			fmt.Println("error : one of the characters or more are missing from the symbols of The asciiArt !")
			os.Exit(1)
		}
	}
}
