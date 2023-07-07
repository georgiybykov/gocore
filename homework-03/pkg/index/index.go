package index

import (
	"strings"

	"gocore/homework-03/pkg/crawler"
)

type Storage struct {
	Data map[string][]int
}

func New() *Storage {
	s := Storage{}
	s.Data = make(map[string][]int)
	return &s
}

func (s *Storage) Append(docs []crawler.Document) map[string][]int {
	for _, doc := range docs {
		for _, lexeme := range parse(doc.Title) {
			if !sliceContains(s.Data[lexeme], doc.ID) {
				s.Data[lexeme] = append(s.Data[lexeme], doc.ID)
			}
		}
	}
	return s.Data
}

func (s *Storage) Search(lexeme string) []int {
	return s.Data[strings.ToLower(lexeme)]
}

func parse(title string) []string {
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
