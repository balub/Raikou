package main

import (
	"os"

	"github.com/raikou/internal/commands"
)

func main() {

	argsWithoutProg := os.Args[1:]

	commands.RunCommand(argsWithoutProg)

}

// rk = list out all hosts defined
// rk -h = list out all the commands and arguments available
