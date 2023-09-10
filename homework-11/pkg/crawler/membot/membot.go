package membot

import (
	"gocore/homework-11/pkg/crawler"
)

// Service - имитация службы поискового робота.
type Service struct{}

// New - конструктор имитации службы поискового робота.
func New() *Service {
	s := Service{}
	return &s
}

// Scan возвращает заранее подготовленный набор данных.
func (s *Service) Scan(url string, depth int) ([]crawler.Document, error) {
	data := []crawler.Document{
		{
			ID:    0,
			URL:   "https://yandex.ru",
			Title: "Яндекс",
		},
		{
			ID:    1,
			URL:   "https://google.ru",
			Title: "Google",
		},
	}

	return data, nil
}
