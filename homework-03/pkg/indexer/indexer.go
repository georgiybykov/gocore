package indexer

import "gocore/homework-03/pkg/crawler"

type Interface interface {
	Add([]crawler.Document)
	Search(string) []int
}
