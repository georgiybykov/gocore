package index

import (
	"strings"

	"gocore/homework-03/pkg/crawler"
)

type Storage struct {
	data map[string][]int
}

func New() *Storage {
	s := Storage{}
	s.data = make(map[string][]int)
	return &s
}

func (s *Storage) Append(docs []crawler.Document) {
	for _, doc := range docs {
		for _, lexeme := range tokenize(doc.Title) {
			if !sliceContains(s.data[lexeme], doc.ID) {
				s.data[lexeme] = append(s.data[lexeme], doc.ID)
			}
		}
	}
}

func (s *Storage) Search(token string) []int {
	return s.data[strings.ToLower(token)]
}

func tokenize(title string) []string {
	lexemes := strings.Split(title, " ")
	for idx := range lexemes {
		lexemes[idx] = strings.ToLower(lexemes[idx])
	}
	return lexemes
}

func sliceContains(ids []int, documentID int) bool {
	for _, id := range ids {
		if id == documentID {
			return true
		}
	}
	return false
}
