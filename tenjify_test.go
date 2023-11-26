package tenjify

import (
	"log"
	"os"
	"testing"
)

const FileName = "sample.png"

func TestTenjify(t *testing.T) {
	reader, err := os.Open(FileName)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	res := Tenjify(reader, 60, 100, false, false)
	t.Log("\n" + res)
}

func TestTenjifyReverse(t *testing.T) {
	reader, err := os.Open(FileName)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	res := Tenjify(reader, 60, 100, true, false)
	t.Log("\n" + res)
}

func TestTenjifyFillBlank(t *testing.T) {
	reader, err := os.Open(FileName)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	res := Tenjify(reader, 60, 100, true, true)
	t.Log("\n" + res)
}
