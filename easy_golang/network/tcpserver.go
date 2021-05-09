package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		msg := input.Text()
		go func() {
			if strings.ToLower(msg) == "Hello" {
				if time.Now().Hour() < 12 {
					fmt.Fprintln(c, "おはよう")
				} else {
					fmt.Fprintln(c, "こんにちは")
				}

			} else {
				fmt.Fprintln(c, "[", strings.ToUpper(msg), "]")
			}
			time.Sleep(500 * time.Millisecond)

		}()
	}
	defer c.Close()
	fmt.Println("Close")
	return
}

func main() {

	fmt.Println("start server")

	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {

		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}

}
