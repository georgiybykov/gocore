package main

import (
	"fmt"

	"gocore/gosearch/pkg/crawler"
	"gocore/gosearch/pkg/crawler/spider"
)

var (
	// word string
	urls  = [2]string{"https://go.dev", "https://golang.org"}
	depth = 1 // 3
)

// flag.StringVar(&word, "s", "", "search query")
// flag.Parse()

// fmt.Println(word)
// fmt.Println(os.Args[1:])

func main() {
	// готовит параметры

	// запускает сканирование по сайтам +
	// объединяет результаты в нумерованный список +

	// если есть параметр поиска фильтрует результаты и возвращает мутированный список

	fmt.Println("Start searching...")

	var documents []crawler.Document

	scanner := spider.New()

	for _, url := range urls {
		result, error := scanner.Scan(url, depth)
		if error != nil {
			fmt.Printf("Error occurs while searching %q", url)
			continue
		}

		documents = append(documents, result...)
	}

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
