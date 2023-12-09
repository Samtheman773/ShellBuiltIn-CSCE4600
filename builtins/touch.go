package builtins

import (
	"fmt"
	"os"
	"time"
)

func TouchFile(args ...string) error {
	switch len(args) {
	case 1:
		fileName := args[0]
		err := touchOrCreateFile(fileName)
		if err != nil {
			return fmt.Errorf("failed to touch file: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("expected exactly one argument (file name)")
	}
}

func touchOrCreateFile(fileName string) error {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		defer file.Close()
		fmt.Println("File created by touch")
		return nil
	}

	currentTime := time.Now().Local()
	err = os.Chtimes(fileName, currentTime, currentTime)
	if err != nil {
		return fmt.Errorf("failed to update file timestamp: %w", err)
	}
	fmt.Println("Timestamp updated by touch")
	return nil
}
