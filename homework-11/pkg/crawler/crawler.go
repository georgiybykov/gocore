package crawler

// Поисковый робот.
// Осуществляет сканирование сайтов.

// Interface определяет контракт поискового робота.
type Interface interface {
	Scan(url string, depth int) ([]Document, error)
	BatchScan(urls []string, depth int, workers int) (<-chan Document, <-chan error)
}

// Document - документ, веб-страница полученная поисковым роботом.
type Document struct {
	ID    int    `json:"id"`
	URL   string `json:"url"`
	Title string `json:"title"`
	Body  string `json:"body,omitempty"`
}
