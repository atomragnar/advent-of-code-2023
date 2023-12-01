package util

import (
	"bufio"
	"os"
	"path/filepath"
)

type BufferProcessor func(*bufio.Reader) error

func processInputPath(inputPath string) (string, error) {

	ip, err := filepath.Abs(inputPath)

	if err != nil {
		return "", err
	}

	return ip, nil

}

func processInputFile(inputPath string) (*os.File, error) {
	ip, err := processInputPath(inputPath)

	if err != nil {
		return nil, err
	}

	inputPath = ip

	inputFile, err := os.Open(inputPath)

	if err != nil {
		return nil, err
	}

	return inputFile, nil
}

func ProcessInput(inputPath string, fn BufferProcessor) error {

	inputFile, err := processInputFile(inputPath)

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
