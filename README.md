
[![Test](https://github.com/fahrizalfarid/read-binary-golang/actions/workflows/test.yml/badge.svg)](https://github.com/fahrizalfarid/read-binary-golang/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/fahrizalfarid/read-binary-golang/branch/main/graph/badge.svg?token=5CD5X1CEFB)](https://codecov.io/gh/fahrizalfarid/read-binary-golang)

## Run
```bash
go run main.go readWrite -i README.md -o text.txt -b 1024
go run main.go readWrite -i img.jpg -o img.png -b 1024
```

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
		_, _ = w.buffer.Write(buffer)
	}

	err := os.WriteFile(output, w.buffer.Bytes(), fs.FileMode(os.O_CREATE))
	if err != nil {
		return err
	}
	return nil
}
```
