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
	word  string
)

// flag.StringVar(&word, "s", "", "search query")
// flag.Parse()

// fmt.Println(word)
// fmt.Println(os.Args[1:])

func main() {
	// готовит параметры +

	// запускает сканирование по сайтам +
	// объединяет результаты в нумерованный список +

	// если есть параметр поиска фильтрует результаты и возвращает мутированный список +

	fmt.Println("Start searching...")

	parseArgs()

	documents := scan(urls, depth)

	printList(documents)
}

func parseArgs() {
	flag.StringVar(&word, "s", "", "Parameter for search")
	flag.Parse()

	if len(word) > 0 {
		word = strings.ToLower(word)
	}
}

func scan(urls [2]string, depth int) []crawler.Document {
	var documents []crawler.Document

	scanner := spider.New()

	for _, url := range urls {
		result, error := scanner.Scan(url, depth)
		if error != nil {
			fmt.Printf("Error occurs while searching %q\n", url)
			continue
		}

		documents = append(documents, result...)
	}

	if len(word) > 0 {
		return filter(documents)
	}

	return documents
}

func filter(documents []crawler.Document) []crawler.Document {
	var filteredDocuments []crawler.Document

	for _, document := range documents {
		title := strings.ToLower(document.Title)

		if strings.Contains(title, word) {
			filteredDocuments = append(filteredDocuments, document)
		}
	}
	return filteredDocuments
}

func printList(documents []crawler.Document) {
	if len(documents) == 0 {
		fmt.Println("Documents not found")
		return
	}

	for index, document := range documents {
		if document.Body != "" {
			fmt.Printf("%d. %q (%q): %q\n", index, document.Title, document.URL, document.Body)
		} else {
			fmt.Printf("%d. %q (%q)\n", index, document.Title, document.URL)
		}
	}
}
