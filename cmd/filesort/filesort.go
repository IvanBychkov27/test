package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	fileName := "cmd/filesort/file/uaAndroid.txt"
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	if len(f) == 0 {
		fmt.Println("data = 0")
		return
	}

	dataFile := strings.Split(string(f), "\n")

	if len(dataFile) == 0 {
		fmt.Println("len(dataFile) == 0")
		return
	}

	sort.Strings(dataFile)

	var fs string
	for _, d := range dataFile {
		fs += d + "\n"
	}

	dataFileSort := []byte(fs)
	fileNameSort := "cmd/filesort/file/uaAndroidSort.txt"

	err = ioutil.WriteFile(fileNameSort, dataFileSort, 0644)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("FileSort: ", fileNameSort)
	fmt.Println("Done")
}
