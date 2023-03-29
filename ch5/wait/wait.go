package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	for _, url := range os.Args[1:] {
		waitForServer(url)
	}
}

func waitForServer(url string) error {
	const timeToWait = 1 * time.Minute
	deadline := time.Now().Add(timeToWait)

	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server no responding (%s);retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeToWait)

}
