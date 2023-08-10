package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handleConn1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, time.Now().Format("2006-01-02 15:04:05"))
}
func main() {
	http.HandleFunc("/", handleConn1)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
