package ascii

import (
	"fmt"
	"os"
)
/* checks whether the input contains an unprintable char or not, if one found it prints an error and exits with status code 1, otherwise it does nothing.*/
func CheckUnprintableChars(s string) {
	for _, c := range s {
		if c >= 32 && c <= 126 {
			continue
		}
		fmt.Println("error one of chars is unprintable !")
		os.Exit(1)
	}
}
