package main

import (
	"fmt"
	"os"
)

func main() {
	content := "this is umesh's file"
	file, err := os.Create("learnFSInGO.txt")
	panicErr(err)
	file.WriteString(content)
	file2, err2 := os.ReadFile("learnFSInGO.txt")
	panicErr(err2)
	fmt.Println(string(file2))
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
