package ascii

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)
/* creates the space to store the  data and filling the space with the correct formated data and returns it to the caller function*/
func ParseFile(banner string) data {
	amap := new(data)
	amap.content = make(map[rune][]string)
	banner += ".txt"
	_, FileName, _, _ := runtime.Caller(0)
	Path := filepath.Join(filepath.Dir(FileName), "../Banner/", banner)
	file, err := ioutil.ReadFile(Path)
	file = file[1 : len(file)-1]
	CheckFileData(string(file))
	if err != nil {
		fmt.Println("Check if ")
		os.Exit(1)
	} else {
		file = []byte(strings.ReplaceAll(string(file), "\r", ""))
		j := 0
		for i := 32; i <= 126; i++ {
			amap.content[rune(i)] = strings.Split(strings.Split(string(file), "\n\n")[j], "\n")
			j++
		}
	}
	return *amap
}
