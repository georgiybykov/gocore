package main

import (
	"flag"
	"fmt"
	"strings"

	"gocore/gosearch/pkg/crawler"
	"gocore/gosearch/pkg/crawler/spider"
)

func main() {
	urls, depth, arg := prepareParams()

	fmt.Println("Start searching...")

	documents := scan(urls, depth, arg)
	render(documents)
}

func scan(urls [2]string, depth int, arg string) (docs []crawler.Document) {
	spider := spider.New()

	for _, url := range urls {
		result, error := spider.Scan(url, depth)
		if error != nil {
			fmt.Printf("An error occurred while searching by URL %q\n", url)
			continue
		}

		docs = append(docs, result...)
	}

	if len(arg) > 0 {
		return filter(docs, arg)
	}
	return docs
}

func filter(documents []crawler.Document, arg string) (docs []crawler.Document) {
	for _, doc := range documents {
		title := strings.ToLower(doc.Title)

		if strings.Contains(title, arg) {
			docs = append(docs, doc)
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

	for idx, doc := range documents {
		if doc.Body != "" {
			fmt.Printf("%d. %q (%q): %q\n", idx, doc.Title, doc.URL, doc.Body)
		} else {
			fmt.Printf("%d. %q (%q)\n", idx, doc.Title, doc.URL)
		}
	}
}

func prepareParams() (urls [2]string, depth int, arg string) {
	urls = [2]string{"https://go.dev", "https://golang.org"}
	depth = 2

	flag.StringVar(&arg, "s", "", "Parameter for search")
	flag.Parse()

	if len(arg) > 0 {
		arg = strings.ToLower(arg)
	}

	return urls, depth, arg
}
