package terminal

import (
	"testing"
)

func Test_HelloWorld(T *testing.T) {
	term := SysTerminal
	term.Print("Hello,", " World")
}

func Test_Read(T *testing.T) {
	term := SysTerminal
	o, e := term.Readln()
	T.Log(o, e)
}
