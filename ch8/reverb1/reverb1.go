package reverb1

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

func handle(c net.Conn) {
	defer c.Close()
	io.Copy(c, c)

}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

func main() {
	listener, _ := net.Listen("tcp", "localhost:8000")
	for {
		conn, _ := listener.Accept()
		go handleConn(conn)
	}
}
