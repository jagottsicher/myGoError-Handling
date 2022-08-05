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
		log.Fatalln(err)
		// panic(err)
	}
}

func foo() {
	fmt.Println("the Function log.Fatalln(err) called os.Exit(1). This deferred function never will be called!")
}
