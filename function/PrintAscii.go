package ascii

import (
	"fmt"
	"strings"
)
/* prints out the input with ascii art symbols */
func PrintAscii(str1 string, amap data) {
	str0 := strings.Split(str1, "\\n")
	if strings.ReplaceAll(str1, "\\n", "") == "" {
		str0 = str0[1:]
	}
	for _, str := range str0 {
		for j := 0; j < 8; j++ {
			if str == "" {
				fmt.Println()
				break
			}
			i := 0
			for i < len(str) {
				SplitedLines := []string(amap.content[rune(str[i])])
				fmt.Print(SplitedLines[j])
				i++
			}
			fmt.Println()
		}
	}
}
