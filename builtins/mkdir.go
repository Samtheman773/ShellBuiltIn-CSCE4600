package builtins

import (
	"fmt"
	"os"
)

func MakeDirectory(args ...string) error {
	switch len(args) {
	case 1:
		err := createDirectory(args[0])
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
		fmt.Println("mkdir is successful")
		return nil
	default:
		return fmt.Errorf("expected exactly one argument (new directory name)")
	}
}

func createDirectory(dirname string) error {
	err := os.Mkdir(dirname, 0750)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("mkdir is unsuccessful")
	}
	return err
}
