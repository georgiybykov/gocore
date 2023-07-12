package repository

import (
	"encoding/json"

	"gocore/homework-05/pkg/crawler"
	"gocore/homework-05/pkg/repository/filestore"
)

const StoreName = "documents.json"

func Filter(lexeme string) ([]crawler.Document, error) {
	file, error := filestore.Fetch(StoreName)
	if error != nil {
		return nil, error
	}
	defer file.Close()

	bytes, error := filestore.Read(file)
	if error != nil {
		return nil, error
	}

	documents := make(map[string][]crawler.Document)
	if error := json.Unmarshal(bytes, &documents); error != nil {
		return nil, error
	}

	return documents[lexeme], nil
}

func Push(documents []crawler.Document, lexeme string) error {
	file, error := filestore.Create(StoreName)
	if error != nil {
		return error
	}
	defer file.Close()

	list := make(map[string][]crawler.Document)
	list[lexeme] = documents

	bytes, error := json.MarshalIndent(list, "", "   ")
	if error != nil {
		return error
	}

	error = filestore.Write(file, bytes)
	return error
}
