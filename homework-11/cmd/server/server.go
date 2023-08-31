package main

import (
	"bufio"
	"fmt"
	"net"
	"sort"
	"time"

	"gocore/homework-11/pkg/crawler"
	"gocore/homework-11/pkg/crawler/spider"
	"gocore/homework-11/pkg/index"
)

const (
	address  = "0.0.0.0:8000"
	protocol = "tcp4"
	depth    = 1
)

func main() {
	resources := [2]string{
		"https://www.practical-go-lessons.com/",
		"https://go.dev",
	}

	fmt.Println("Start searching...")

	documents := scan(resources, depth)

	fmt.Println("Searching is finished")

	sort.SliceStable(documents, func(i, j int) bool {
		return documents[i].ID < documents[j].ID
	})

	fmt.Println("Start indexing...")

	indexer := index.New()
	indexer.Append(documents)

	fmt.Println("Indexed")
	fmt.Println("Ready for requests")

	// render(documents)

	// from here

	listener, err := net.Listen(protocol, address)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handler(conn, indexer, documents)
	}
}

func handler(conn net.Conn, indexer *index.Storage, documents []crawler.Document) {
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(time.Second * 60))

	r := bufio.NewReader(conn)

	for {
		lexeme, _, err := r.ReadLine()
		if err != nil {
			return
		}

		if len(lexeme) == 0 {
			fmt.Println("The lexeme to search for not found.")
			conn.Write([]byte("The lexeme to search for not found."))
			return
		}

		indices := indexer.Search(string(lexeme))
		documents = filter(documents, indices)

		// render(documents)

		if len(documents) == 0 {
			fmt.Println("Documents not found.")
			conn.Write([]byte("Documents not found."))
			return
		}

		fmt.Println("Documets:")
		conn.Write([]byte("Documets:\n"))

		for _, doc := range documents {
			if doc.Body != "" {
				fmt.Printf("%d. %q (%q): %q\n", doc.ID, doc.Title, doc.URL, doc.Body)

				s := fmt.Sprintf("%d. %q (%q): %q\n", doc.ID, doc.Title, doc.URL, doc.Body)

				conn.Write([]byte(s))
			} else {
				fmt.Printf("%d. %q (%q)\n", doc.ID, doc.Title, doc.URL)

				s := fmt.Sprintf("%d. %q (%q)\n", doc.ID, doc.Title, doc.URL)

				conn.Write([]byte(s))
			}
		}
	}
}

func scan(resources [2]string, depth int) (docs []crawler.Document) {
	spider := spider.New()

	for _, resource := range resources {
		doc, err := spider.Scan(resource, depth)
		if err != nil {
			fmt.Printf("An error occurred while searching by resource %q\n", resource)
			continue
		}

		docs = append(docs, doc...)
	}

	for idx := range docs {
		docs[idx].ID = idx + 1
	}

	return docs
}

func filter(documents []crawler.Document, indices []int) (docs []crawler.Document) {
	for _, id := range indices {
		idx := sort.Search(len(documents), func(idx int) bool {
			return documents[idx].ID >= id
		})

		if idx <= len(documents) && documents[idx].ID == id {
			docs = append(docs, documents[idx])
		}
	}
	return docs
}

// func render(documents []crawler.Document) {
// 	if len(documents) == 0 {
// 		fmt.Println("Documents not found")
// 		return
// 	}

// 	fmt.Println("Documets:")

// 	for _, doc := range documents {
// 		if doc.Body != "" {
// 			fmt.Printf("%d. %q (%q): %q\n", doc.ID, doc.Title, doc.URL, doc.Body)
// 		} else {
// 			fmt.Printf("%d. %q (%q)\n", doc.ID, doc.Title, doc.URL)
// 		}
// 	}
// }
