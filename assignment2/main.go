package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f := readFile(os.Args[1])
	defer f.Close()

	//lw := logWriter{}
	_, err := io.Copy(os.Stdout, f)
	if err != nil {
		return
	}
}

func readFile(filename string) *os.File {
	f, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		os.Exit(1)
	}
	return f
}
