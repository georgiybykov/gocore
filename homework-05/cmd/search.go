package main

import (
	"flag"
	"fmt"
	"sort"

	"gocore/homework-05/pkg/crawler"
	"gocore/homework-05/pkg/crawler/spider"
	"gocore/homework-05/pkg/index"
	"gocore/homework-05/pkg/repository"
)

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

	documents, _ := repository.Filter(*lexeme)
	if documents != nil {
		render(documents)
		return
	}

	documents = scan(urls, depth)

	sort.SliceStable(documents, func(i, j int) bool {
		return documents[i].ID < documents[j].ID
	})

	indexer := index.New()
	indexer.Append(documents)
	indices := indexer.Search(*lexeme)
	documents = filter(documents, indices)

	error := repository.Push(documents, *lexeme)
	if error != nil {
		fmt.Printf("Error writing to the Documents repository:\n    - %v\n", error)
	}

	render(documents)
}

func scan(urls [2]string, depth int) (docs []crawler.Document) {
	spider := spider.New()

	for _, url := range urls {
		doc, error := spider.Scan(url, depth)
		if error != nil {
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
