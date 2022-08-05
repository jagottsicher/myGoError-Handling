package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("a-file.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

	ourContent := strings.NewReader("He's a family guuuuuuuy!\n")

	io.Copy(f, ourContent)
}
