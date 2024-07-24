package ascii

import "strings"

/* protects for error cases */

func Protect(arg1 string, banner *string, n int) {
	CheckUnprintableChars(arg1)
	if n == 1 {
		return
	} else {
		if *banner == "shadow.txt" || *banner == "standard.txt" || *banner == "thinkertoy.txt" {
			*banner = strings.Split(*banner, ".")[0]
		}
	}
}
