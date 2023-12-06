package util

import (
	"bufio"
	"os"
	"path/filepath"
)

type BufferProcessor func(*bufio.Reader) error

func ProcessInput(inputPath string, fn BufferProcessor) error {
	ip, err := filepath.Abs(inputPath)

	if err != nil {
		return err
	}

	inputPath = ip

	inputFile, err := os.Open(inputPath)

	if err != nil {
		return err
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {

		}
	}(inputFile)

	reader := bufio.NewReader(inputFile)

	err = fn(reader)

	if err != nil {
		return err
	}

	return nil

}
