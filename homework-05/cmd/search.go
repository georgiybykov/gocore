package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"

	"gocore/homework-05/pkg/crawler"
	"gocore/homework-05/pkg/crawler/spider"
	"gocore/homework-05/pkg/index"
	"gocore/homework-05/pkg/repository"
)

const FilePath = "./homework-05/pkg/repository/documents.json"

func main() {
	urls := [2]string{"https://go.dev", "https://www.practical-go-lessons.com/"}
	depth := 2

	lexeme := flag.String("s", "", "Parameter for search")
	flag.Parse()

	if len(*lexeme) == 0 {
		fmt.Println("The lexeme to search for not found. Try: `go run [command] -s [lexeme]`")
		return
	}

	fmt.Println("Start searching...")

	var documents []crawler.Document

	if _, err := os.Stat(FilePath); err == nil {
		file, err := os.Open(FilePath)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		defer file.Close()

		bytes, err := repository.Read(file)
		if err != nil {
			fmt.Printf("Error reading from the Documents repository:\n    - %v\n", err)
			return
		}

		if err := json.Unmarshal(bytes, &documents); err != nil {
			fmt.Println("Error: ", err)
			return
		}
	} else if errors.Is(err, os.ErrNotExist) {
		documents = scan(urls, depth)

		sort.SliceStable(documents, func(i, j int) bool {
			return documents[i].ID < documents[j].ID
		})

		file, err := os.Create(FilePath)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		defer file.Close()

		bytes, err := json.MarshalIndent(documents, "", "   ")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		err = repository.Write(file, bytes)
		if err != nil {
			fmt.Printf("Error writing to the Documents repository:\n    - %v\n", err)
			return
		}
	} else {
		fmt.Println("Error: ", err)
		return
	}

	indexer := index.New()
	indexer.Append(documents)
	indices := indexer.Search(*lexeme)
	documents = filter(documents, indices)

	render(documents)
}

func scan(urls [2]string, depth int) (docs []crawler.Document) {
	spider := spider.New()

	for _, url := range urls {
		doc, err := spider.Scan(url, depth)
		if err != nil {
			fmt.Printf("An error occurred while searching by URL %q\n", url)
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

func render(documents []crawler.Document) {
	if len(documents) == 0 {
		fmt.Println("Documents not found")
		return
	}

	fmt.Println("Documets:")

	for _, doc := range documents {
		if doc.Body != "" {
			fmt.Printf("%d. %q (%q): %q\n", doc.ID, doc.Title, doc.URL, doc.Body)
		} else {
			fmt.Printf("%d. %q (%q)\n", doc.ID, doc.Title, doc.URL)
		}
	}
}
