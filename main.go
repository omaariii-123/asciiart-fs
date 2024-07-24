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


func CheckUnprintableChars(s string) {
	for _, c := range s {
		if c >= 32 && c <= 126 {
			continue
		}
		fmt.Println("error one of chars is unprintable !")
		os.Exit(1)
	}
}

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

func PrintAscii(str1 string, amap data) {
  str0 := strings.Split(str1, "\\n")
  if strings.ReplaceAll(str1, "\\n", "") == ""{
    str0 = str0[1:]
  }
  for _, str := range str0 {
    for j := 0; j < 8; j++ {
      if str == "" {
        fmt.Println()
        break
      }
        i := 0
	    	for i < len(str){
		    	SplitedLines := []string(amap.content[rune(str[i])])
		    	fmt.Print(SplitedLines[j])
		    	i++
			}
  			fmt.Println()
    	}
  }
}

func ParseFile(banner string) data{
	amap := new(data)
	amap.content = make(map[rune][]string)
	banner += ".txt"
	file, err := ioutil.ReadFile(banner)
	if err != nil {
		fmt.Println("error reading the banner file !")
		os.Exit(1)
	}else { 
      file = []byte(strings.ReplaceAll(string(file), "\r", ""))[1:]
		j := 0
		for i := 32; i <= 126; i++ {
			amap.content[rune(i)] = strings.Split(strings.Split(string(file), "\n\n")[j],"\n")
			j++
		}
	}
	return *amap
}
func main() {
	if len(os.Args) == 3 {
		if (os.Args[1] != "" && os.Args[2] != "") {
			Protect(os.Args[1], &os.Args[2], 2)
			amap := ParseFile(os.Args[2])
			PrintAscii(os.Args[1], amap)
		} 	
	}else if len(os.Args) == 2 {
				Protect(os.Args[1], nil, 1)
				amap := ParseFile("standard")
				PrintAscii(os.Args[1], amap)
		}

}
