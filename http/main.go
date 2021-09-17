package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logW struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	lw := logW{}
	io.Copy(lw, resp.Body)
}

//log w implementing now the Write interface
func (logW) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("custom implementation here")
	return len(bs), nil
}
