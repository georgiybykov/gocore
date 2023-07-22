package repository

import (
	"bufio"
	"io"
)

func Read(reader io.Reader) ([]byte, error) {
	scanner := bufio.NewScanner(reader)

	var bytes []byte

	for scanner.Scan() {
		bytes = append(bytes, []byte(scanner.Text())...)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return bytes, nil
}

func Write(writer io.Writer, bytes []byte) error {
	_, err := writer.Write(append(bytes, '\n'))
	return err
}
