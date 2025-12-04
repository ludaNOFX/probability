package randomvariablemodel

import "os"

func GetStdin() *os.File {
	return os.Stdin
}
