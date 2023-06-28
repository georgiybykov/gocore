package main

import (
	"flag"
	"fmt"
	"strings"

	"gocore/gosearch/pkg/crawler"
	"gocore/gosearch/pkg/crawler/spider"
)

var (
	urls  = [2]string{"https://go.dev", "https://golang.org"}
	depth = 1 // 3
)

func main() {
	fmt.Println("Start searching...")

	argument := extractArgument()
	documents := scan(urls, depth, argument)
	render(documents)
}

func scan(urls [2]string, depth int, argument string) (documents []crawler.Document) {
	scanner := spider.New()

	for _, url := range urls {
		result, error := scanner.Scan(url, depth)
		if error != nil {
			fmt.Printf("An error occurred while searching by URL %q\n", url)
			continue
		}

		documents = append(documents, result...)
	}

	if len(argument) > 0 {
		return filter(documents, argument)
	}
	return documents
}

func filter(documents []crawler.Document, argument string) (filteredDocuments []crawler.Document) {
	for _, document := range documents {
		title := strings.ToLower(document.Title)

		if strings.Contains(title, argument) {
			filteredDocuments = append(filteredDocuments, document)
		}
	}
	return filteredDocuments
}

func render(documents []crawler.Document) {
	if len(documents) == 0 {
		fmt.Println("Documents not found")
		return
	}

	fmt.Println("Documets:")

	for index, document := range documents {
		if document.Body != "" {
			fmt.Printf("%d. %q (%q): %q\n", index, document.Title, document.URL, document.Body)
		} else {
			fmt.Printf("%d. %q (%q)\n", index, document.Title, document.URL)
		}
	}
}

func extractArgument() (argument string) {
	flag.StringVar(&argument, "s", "", "Parameter for search")
	flag.Parse()

	if len(argument) > 0 {
		argument = strings.ToLower(argument)
	}
	return argument
}
