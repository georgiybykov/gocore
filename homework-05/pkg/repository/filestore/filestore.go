package filestore

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
)

const FileStorePath = "./homework-05/pkg/repository/filestore/"

func Fetch(name string) (*os.File, error) {
	path, error := filepath.Abs(FileStorePath + name)
	if error != nil {
		return nil, error
	}

	_, error = os.Stat(path)
	if error != nil {
		return nil, error
	}

	file, error := os.Open(path)
	if error != nil {
		return nil, error
	}

	return file, nil
}

func Create(name string) (*os.File, error) {
	path, error := filepath.Abs(FileStorePath + name)
	if error != nil {
		return nil, error
	}

	file, error := os.Create(path)
	if error != nil {
		return nil, error
	}

	return file, error
}

func Read(reader io.Reader) ([]byte, error) {
	scanner := bufio.NewScanner(reader)

	var bytes []byte

	for scanner.Scan() {
		bytes = append(bytes, []byte(scanner.Text())...)
	}

	if error := scanner.Err(); error != nil {
		return nil, error
	}

	return bytes, nil
}

func Write(writer io.Writer, bytes []byte) error {
	_, error := writer.Write(append(bytes, '\n'))
	return error
}
