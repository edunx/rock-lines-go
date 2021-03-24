package lines

import (
	"bufio"
	"github.com/edunx/lua"
	"os"
)

type Lines struct {
	lua.Super

	filename string
	fd       *os.File
	scanner  *bufio.Scanner
}
