package lines

import (
	"bufio"
	"os"
)

type Lines struct {
	filename string
	fd       *os.File
	scanner  *bufio.Scanner
}
