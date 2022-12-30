package usecase

import (
	"io"
	"log"
	"os"
)

type (
	read struct{}
	Read interface {
		ReadFile(filepath string, bufferSize int) (<-chan []byte, error)
	}
)

func NewReader() Read {
	return &read{}
}

func (r *read) ReadFile(filepath string, bufferSize int) (<-chan []byte, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	bufferChan := make(chan []byte, bufferSize)

	go func() {
		defer f.Close()
		defer close(bufferChan)

		buffer := make([]byte, bufferSize)
		for {
			n, err := f.Read(buffer)
			if err == io.EOF {
				log.Println("EOF")
				break
			}

			bufferChan <- buffer[:n]
		}
	}()

	return bufferChan, nil
}
