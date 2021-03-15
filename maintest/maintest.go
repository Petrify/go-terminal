package main

import (
	"log"

	"github.com/Petrify/go-terminal"
)

func main() {
	t := terminal.SysTerminal
	t.Println("Type yo thing here:")
	log.Println(t.Readln())
}
