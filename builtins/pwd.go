package builtins

import (
	//"errors"
	"fmt"
	"io"
	"os"
	//"strings"
)

var (
// ErrInvalidArgCount = errors.New("invalid argument count")
// HomeDir, _         = os.UserHomeDir()
)

func PrintWorkingDirectory(w io.Writer, args ...string) error {

	//var L = false
	//var P = false

	for i := 0; i < len(args); i++ {
		if args[i] == "-L" || args[i] == "-LP" || args[i] == "-PL" {
			//L = true
		}
		if args[i] == "-P" || args[i] == "-LP" || args[i] == "-PL" {
			//P = true
		}
	}

	//Probably handle L and P differently

	dir, wdErr := os.Getwd()
	if wdErr != nil {
		return wdErr
	}
	_, err := fmt.Fprintln(w, dir)

	return err
}
