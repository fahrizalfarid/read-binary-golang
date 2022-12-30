## Run
`go run main.go readWrite -i README.md -o text.txt -b 1024`

## src/usecase/read.go
```go
package usecase

import (
	"io"
	"log"
	"os"
)

type (
	read struct{}
	Read interface {
		ReadFile(filepath string, bufferSize int) <-chan []byte
	}
)

func NewReader() Read {
	return &read{}
}

func (r *read) ReadFile(filepath string, bufferSize int) <-chan []byte {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
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

	return bufferChan
}
```

## src/usecase/write.go
```go
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
		_, err := w.buffer.Write(buffer)
		if err != nil {
			return err
		}
	}

	err := os.WriteFile(output, w.buffer.Bytes(), fs.FileMode(os.O_CREATE))
	if err != nil {
		return err
	}
	return nil
}
```