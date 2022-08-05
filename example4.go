package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("a-file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	readContent, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(readContent))
}