package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.google.com")
	//c := resp.Header["Content-Type"]
	ct := strings.Split(resp.Header["Content-Type"][0], ";")
	fmt.Println(ct[0])

	lw := logWriter{}
	_, err = io.Copy(lw, resp.Body)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	switch ct[0] {
	case "text/html":
		fmt.Println("content type is text/html")
	case "linux":
		fmt.Println("Go runs on Linux.")
	default:
		fmt.Printf("Go runs on ")
	}
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
