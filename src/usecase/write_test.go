package usecase

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteImgFile(t *testing.T) {
	defer func() {
		os.Remove("../../img.png")
	}()

	reader := NewReader()
	writer := NewWriter()

	chanBin := reader.ReadFile("../../img.jpg", 1024)
	err := writer.WriteToFile(chanBin, "../../img.png")
	assert.Nil(t, err, "file already exists")
}

func TestWriteTxtFile(t *testing.T) {
	defer func() {
		os.Remove("../../README.csv")
	}()

	reader := NewReader()
	writer := NewWriter()

	chanBin := reader.ReadFile("../../README.md", 1024)
	err := writer.WriteToFile(chanBin, "../../README.csv")
	assert.Nil(t, err, "file already exists")
}

func BenchmarkWriteImgFile(t *testing.B) {
	defer func() {
		os.Remove("../../benchmark.png")
	}()

	reader := NewReader()
	writer := NewWriter()

	chanBin := reader.ReadFile("../../img.jpg", 1024)
	_ = writer.WriteToFile(chanBin, "../../benchmark.png")
}

func BenchmarkWriteVideoFile(t *testing.B) {
	defer func() {
		os.Remove("../../benchmark.mkv")
	}()

	reader := NewReader()
	writer := NewWriter()

	chanBin := reader.ReadFile("../../kambing.mp4", 1000000)
	_ = writer.WriteToFile(chanBin, "../../benchmark.mkv")
}

func BenchmarkWriteTxtFile(t *testing.B) {
	defer func() {
		os.Remove("../../benchmark.csv")
	}()

	reader := NewReader()
	writer := NewWriter()

	chanBin := reader.ReadFile("../../README.md", 1024)
	_ = writer.WriteToFile(chanBin, "../../benchmark.csv")
}
