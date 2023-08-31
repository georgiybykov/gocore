package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial(protocol, address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	defer fmt.Println("Connection closed")

	terminal := bufio.NewReader(os.Stdin)
	server := bufio.NewReader(conn)

	for {
		fmt.Println("Enter a lexeme to search for:")

		lexeme, _, err := terminal.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Searching for: %q\n", string(lexeme))
		fmt.Fprintf(conn, "%s\n", string(lexeme))

	HANDLER:
		for {
			bytes, _, err := server.ReadLine()
			if err != nil {
				log.Fatal(err)
			}

			response := strings.TrimSpace(string(bytes))
			if len(response) == 0 {
				break HANDLER
			}

			fmt.Println(response)
		}
	}
}
