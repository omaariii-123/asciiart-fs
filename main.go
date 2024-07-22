package main

import (
		"fmt"
		"os"
		"io/ioutil"
		"strings"
	)

type data struct {
	content	map[rune][]string
}

func PrintAscii(str string, amap data, f string) {
		for j := 0; j < 8; j++ {
			i := 0
			for i < len(str){
				SplitedLines := []string(amap.content[rune(str[i])])
				fmt.Print(SplitedLines[j])
				i++
			}
				fmt.Println()
		}
}

func ParseFile(banner string) data {
	amap := new(data)
	amap.content = make(map[rune][]string)
	banner += ".txt"
	file, err := ioutil.ReadFile(banner)
	j := 0
	for i := 32; i < 126; i++ {
		amap.content[rune(i)] = strings.Split(strings.Split(string(file), "\n\n")[j],"\n")
		j++
	}
	if err != nil {
		fmt.Println("error reading the banner file !")
	} 
	return *amap
}
func main() {
	if len(os.Args) == 3 {
		if (os.Args[1] != "" && os.Args[2] != "") {
				amap := ParseFile(os.Args[2])
		//		fmt.Print(amap.content['f'])
			if os.Args[2] == "standard" {
				PrintAscii(os.Args[1], amap, "st")
			} else if os.Args[2] == "shadow" {
				PrintAscii(os.Args[1], amap, "sh")
			} else if os.Args[2] == "thinkertoy" {
				PrintAscii(os.Args[1], amap, "th")
			}
		}
	}
}
