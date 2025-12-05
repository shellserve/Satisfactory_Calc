// Package io handles any and all file operations
// This package will likely also handle logging mechnaisms and user input later on
package io

import (
	"errors"
	"log"
	"os"
)

func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !errors.Is(err, os.ErrNotExist)
}

func WriteStringToFile(filepath string, data string) (bool, error) {
	// will truncate if used on existing file.
	f, err := os.Create(filepath)
	if err != nil {
		return false, err
	}

	// defer file close
	defer func() {
		closeErr := f.Close()
		if closeErr != nil {
			log.Fatalf("Error in closing file hook: %v", closeErr)
		}
	}()

	_, err = f.WriteString(data)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ReadFromFile(filepath string) ([]byte, error) {
	if contentBytes, err := os.ReadFile(filepath); err != nil {
		return nil, err
	} else {
		return contentBytes, nil
	}
}
