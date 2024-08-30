package util

import (
	"os"
)

func ValidateArgOrFullPath(args []string, path string) bool {
    if path[len(path)-1] != os.PathSeparator && len(args) < 1 { return false }
		return true
}
