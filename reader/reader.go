package reader

import (
	"bufio"
	"os"
)

func Read(fileName string) (*bufio.Reader, error) {
	stream, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return bufio.NewReader(stream), nil
}
