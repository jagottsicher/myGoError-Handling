package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("a-file.txt")
	if err != nil {
		// fmt.Println("An error occured:", err)
		// log.Println("An error occured:", err)
		// log.Fatalln(err)
		log.Panicln(err)
		// panic(err)
	}
}

func foo() {
	fmt.Println("The Function log.Panicln(err) ran without calling os.Exit(1). This deferred function will be called!")
}
