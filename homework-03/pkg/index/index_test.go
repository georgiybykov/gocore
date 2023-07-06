package index

import (
	"reflect"
	"testing"

	"gocore/homework-03/pkg/crawler"
)

func TestAppend(t *testing.T) {
	documents := []crawler.Document{
		{
			ID:    1,
			Title: "Two words",
		},
		{
			ID:    3,
			Title: "Three words here",
		},
	}

	storage := New()
	storage.Append(documents)

	result := len(storage.Data)
	expectation := 4
	if result != expectation {
		t.Errorf("expected '%d' but got '%d'", expectation, result)
	}
}

func TestSearch(t *testing.T) {
	documents := []crawler.Document{
		{
			ID:    1,
			Title: "Two words",
		},
		{
			ID:    5,
			Title: "Three words here",
		},
		{
			ID:    15,
			Title: "Four words right here",
		},
	}

	storage := New()
	storage.Append(documents)

	presets := []struct {
		context     string
		lexeme      string
		expectation []int
	}{
		{
			context:     "When the expected lexeme is present in documents",
			lexeme:      "WorDs",
			expectation: []int{1, 5, 15},
		},
		{
			context:     "When the expected lexeme is present in document",
			lexeme:      "two",
			expectation: []int{1},
		},
		{
			context:     "When the expected lexeme is not found",
			lexeme:      "Undefined",
			expectation: nil,
		},
		{
			context:     "When the expected lexeme is more than one word",
			lexeme:      "words right",
			expectation: nil,
		},
	}

	for _, preset := range presets {
		t.Run(preset.context, func(t *testing.T) {
			result := storage.Search(preset.lexeme)
			if !reflect.DeepEqual(result, preset.expectation) {
				t.Errorf("expected '%d' but got '%d'", preset.expectation, result)
			}
		})
	}
}
