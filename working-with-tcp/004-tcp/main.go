package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if i == 0 {
			method, url := strings.Fields(line)[0], strings.Fields(line)[1]
			fmt.Println("***METHOD", method)
			fmt.Println("***URL", url)
			route(conn, url)
		}
		if line == "" {
			break
		}
		i++
	}
}

func route(conn net.Conn, url string) {
	var body string

	switch url {
	case "/":
		body = `
    <!DOCTYPE html>
    <html lang="en">
      <head>
        <meta charset="UTF-8">
        <title>HTTP</title>
      </head>
      <body>
        <header>
          <nav>
            <ul>
              <li><a href="/">Home</a></li>
              <li><a href="/about">About</a></li>
              <li><a href="contact">Contact</a></li>
            </ul>
          </nav
        </header>
        <main>Home Page</main>
      </body>
    </html>
  `
	case "/about":
		body = `
    <!DOCTYPE html>
    <html lang="en">
      <head>
        <meta charset="UTF-8">
        <title>HTTP</title>
      </head>
      <body>
        <header>
          <nav>
            <ul>
              <li><a href="/">Home</a></li>
              <li><a href="/about">About</a></li>
              <li><a href="contact">Contact</a></li>
            </ul>
          </nav
        </header>
        <main>About Page</main>
      </body>
    </html>
    `
	case "/contact":
		body = `
    <!DOCTYPE html>
    <html lang="en">
      <head>
        <meta charset="UTF-8">
        <title>HTTP</title>
      </head>
      <body>
        <header>
          <nav>
            <ul>
              <li><a href="/">Home</a></li>
              <li><a href="/about">About</a></li>
              <li><a href="contact">Contact</a></li>
            </ul>
          </nav
        </header>
        <main>Contact Page</main>
      </body>
    </html>
    `
	}

	respond(conn, body)
}

func respond(conn net.Conn, body string) {
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
