package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex

var count int

func main() {
	http.HandleFunc("/", handler1)
	http.HandleFunc("/count/", counter)

	lissajousHandler := func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		cycles := r.Form.Get("cycles")
		cyclesInt, _ := strconv.ParseFloat(cycles, 32)
		lissajous(w, cyclesInt)
	}
	http.HandleFunc("/lissajous", lissajousHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
