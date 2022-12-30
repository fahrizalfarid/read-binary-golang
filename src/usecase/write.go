package usecase

import (
	"bytes"
	"io/fs"
	"os"
)

type (
	write struct {
		buffer *bytes.Buffer
	}
	Write interface {
		WriteToFile(bufferChan <-chan []byte, output string) error
	}
)

func NewWriter() Write {
	return &write{
		buffer: &bytes.Buffer{},
	}
}

func (w *write) WriteToFile(bufferChan <-chan []byte, output string) error {
	defer func() {
		w.buffer.Reset()
	}()

	for buffer := range bufferChan {
		_, _ = w.buffer.Write(buffer)
	}

	err := os.WriteFile(output, w.buffer.Bytes(), fs.FileMode(os.O_CREATE))
	if err != nil {
		return err
	}
	return nil
}
