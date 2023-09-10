package main

import (
	"fmt"
	"log"
	"sort"

	"gocore/homework-11/pkg/crawler"
	"gocore/homework-11/pkg/crawler/spider"
	"gocore/homework-11/pkg/index"
	"gocore/homework-11/pkg/netsrv/server"
)

const (
	address  = "0.0.0.0:8000"
	protocol = "tcp4"
	depth    = 2
)

func main() {
	resources := [2]string{
		"https://www.practical-go-lessons.com/",
		"https://go.dev",
	}

	documents := scan(resources, depth)

	sort.SliceStable(documents, func(i, j int) bool {
		return documents[i].ID < documents[j].ID
	})

	indexer := index.New()
	indexer.Append(documents)

	fmt.Println("Indexed")

	server, err := server.New(search, documents, indexer)
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	server.Run()
}

func search(lexeme string, documents []crawler.Document, indexer *index.Storage) (bytes []byte) {
	if len(lexeme) == 0 {
		return []byte("The lexeme to search for not found.")
	}

	indices := indexer.Search(lexeme)
	documents = filter(documents, indices)

	if len(documents) == 0 {
		return []byte("Documents not found.")
	}

	bytes = append(bytes, []byte("Documets:\n")...)

	for _, doc := range documents {
		bytes = append(bytes, []byte(fmt.Sprintf("%d. %q (%q): %q\n", doc.ID, doc.Title, doc.URL, doc.Body))...)
	}
	return bytes
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
