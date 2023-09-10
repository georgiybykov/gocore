package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gocore/homework-11/pkg/netsrv/client"
)

func main() {
	client, err := client.New()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	defer fmt.Println("Connection closed")

	terminal := bufio.NewReader(os.Stdin)
	// server := bufio.NewReader(client.Conn)

	for {
		fmt.Println("Enter a lexeme to search for:")

		lexeme, _, err := terminal.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Searching for: %q\n", string(lexeme))

		client.Search(string(lexeme))

		// 	fmt.Fprintf(conn, "%s\n", string(lexeme))

		// HANDLER:
		// 	for {
		// 		bytes, _, err := server.ReadLine()
		// 		if err != nil {
		// 			log.Fatal(err)
		// 		}

		// 		response := strings.TrimSpace(string(bytes))
		// 		if len(response) == 0 {
		// 			break HANDLER
		// 		}

		// 		fmt.Println(response)
		// 	}
	}
}
