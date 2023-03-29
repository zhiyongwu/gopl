package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//b, err := io.ReadAll(resp.Body)

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading : %v\n", err)
			os.Exit(1)
		}

		resp.Body.Close()
		fmt.Printf("%s", b)
		//fmt.Println(written)
		//if nil != err {
		//	fmt.Fprintf(os.Stderr, "fetch reading : %v\n", err)
		//	os.Exit(1)
		//}
		//fmt.Printf("%s", b)
	}
}
