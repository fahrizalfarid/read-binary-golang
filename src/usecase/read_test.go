package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestReadImgFile(t *testing.T) {
	dataChanActual := [][]byte{}
	dataChanTested := [][]byte{}

	assert := assert.New(t)

	readerTested := NewReader()
	readerActual := NewReader()

	chanActual := readerActual.ReadFile("../../img.jpg", 1024)
	chanTested := readerTested.ReadFile("../../img.jpg", 1024)

	for de := range chanActual {
		dataChanActual = append(dataChanActual, de)
	}

	for dt := range chanTested {
		dataChanTested = append(dataChanTested, dt)
	}

	assert.Equal(dataChanTested, dataChanActual)
}

func TestReadTxtFile(t *testing.T) {
	dataChanActual := [][]byte{}
	dataChanTested := [][]byte{}

	assert := assert.New(t)

	readerTested := NewReader()
	readerActual := NewReader()

	chanActual := readerActual.ReadFile("../../README.md", 1024)
	chanTested := readerTested.ReadFile("../../README.md", 1024)

	for de := range chanActual {
		dataChanActual = append(dataChanActual, de)
	}

	for dt := range chanTested {
		dataChanTested = append(dataChanTested, dt)
	}

	assert.Equal(dataChanTested, dataChanActual)
}

func BenchmarkReadTxtFile(t *testing.B) {
	dataBin := [][]byte{}
	defer func() {
		dataBin = nil
	}()

	readerTested := NewReader()

	chanBin := readerTested.ReadFile("../../README.md", 1024)
	for d := range chanBin {
		dataBin = append(dataBin, d)
	}
}

func BenchmarkReadImgFile(t *testing.B) {
	dataBin := [][]byte{}
	defer func() {
		dataBin = nil
	}()

	readerTested := NewReader()

	chanBin := readerTested.ReadFile("../../img.jpg", 1024)
	for d := range chanBin {
		dataBin = append(dataBin, d)
	}
}

func BenchmarkReadVideoFile(t *testing.B) {
	dataBin := [][]byte{}
	defer func() {
		dataBin = nil
	}()

	readerTested := NewReader()

	chanBin := readerTested.ReadFile("../../kambing.mp4", 1000000)
	for d := range chanBin {
		dataBin = append(dataBin, d)
	}
}
