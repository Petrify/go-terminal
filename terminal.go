package terminal

import (
	"bufio"
	"io"
	"log"
)

var ()

type Terminal struct {
	scanner *bufio.Scanner
	*log.Logger
}

func New(in io.Reader, out io.Writer) *Terminal {
	return &Terminal{
		bufio.NewScanner(in),
		log.New(out, "", log.Ltime),
	}
}

// func (t *Terminal) SubLog(label string) *log.Logger {
// 	if label == "" {
// 		return log.New(t, fmt.Sprintf("", label), 0)
// 	}
// 	return log.New(t, fmt.Sprintf("[%s]", label), 0)
// }

func (t *Terminal) Readln() (s string, e error) {
	if !t.scanner.Scan() {
		return "", t.scanner.Err()
	}

	s = t.scanner.Text()
	return s, t.scanner.Err()
}
