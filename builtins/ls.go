package builtins

import (
	"fmt"
	"os"
)

func ListFilesDirectory(args ...string) error {
	switch len(args) {
	case 0:
		return listCurrentDirectory()
	case 1:
		return listDirectoryContent(args[0])
	default:
		return fmt.Errorf("expected zero or one argument (ls)")
	}
}

func listCurrentDirectory() error {
	files, err := os.ReadDir(".")
	if err != nil {
		return fmt.Errorf("failed to read current directory: %v", err)
	}

	printFileNames(files)
	return nil
}

func listDirectoryContent(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read directory '%s': %v", dir, err)
	}

	printFileNames(files)
	return nil
}

func printFileNames(files []os.DirEntry) {
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
