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

	got := len(storage.Data)
	want := 4
	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
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
		context string
		lexeme  string
		want    []int
	}{
		{
			context: "When the expected lexeme is present in documents",
			lexeme:  "WorDs",
			want:    []int{1, 5, 15},
		},
		{
			context: "When the expected lexeme is present in document",
			lexeme:  "two",
			want:    []int{1},
		},
		{
			context: "When the expected lexeme is not found",
			lexeme:  "Undefined",
			want:    nil,
		},
		{
			context: "When the expected lexeme is more than one word",
			lexeme:  "words right",
			want:    nil,
		},
	}

	for _, preset := range presets {
		t.Run(preset.context, func(t *testing.T) {
			got := storage.Search(preset.lexeme)
			if !reflect.DeepEqual(got, preset.want) {
				t.Errorf("expected '%d' but got '%d'", preset.want, got)
			}
		})
	}
}
