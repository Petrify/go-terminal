package terminal

import "os"

var SysTerminal *Terminal

func init() {
	SysTerminal = New(os.Stdin, os.Stdout)
}
