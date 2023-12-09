package builtins

import (
	"fmt"
	"os"
)

func OpenFileDirectory(args ...string) error {
	switch len(args) {
	case 1:
		err := openAndCloseFile(args[0])
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		fmt.Println("Open is successful")
		return nil

	default:
		return fmt.Errorf("expected exactly one argument (file name)")
	}
}

func openAndCloseFile(filename string) error {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	return nil
}
