package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, _ := os.Create(local)
	n, _ = io.Copy(f, resp.Body)
	return local, n, err
}

func main() {
	fmt.Println(fetch(os.Args[1]))
}
