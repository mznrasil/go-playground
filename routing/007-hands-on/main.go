package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		defer conn.Close()

		var method, url, body string

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
			method = strings.Fields(line)[0]
			url = strings.Fields(line)[1]

			switch {
			case method == "GET" && url == "/":
				body = index()
			case method == "GET" && url == "/apply":
				body = apply()
			}
			break
		}

		io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprintf(conn, "Content-Type: text/html\r\n")
		io.WriteString(conn, "\r\n")
		io.WriteString(conn, body)
	}
}

func index() string {
	body := "<!doctype html><html lang=''en'><head><meta charset='utf-8' /><title>Index</title></head><body>Index Route</body></html>"
	return body
}

func apply() string {
	body := "<!doctype html><html lang=''en'><head><meta charset='utf-8' /><title>Index</title></head><body>Apply Route</body></html>"
	return body
}
