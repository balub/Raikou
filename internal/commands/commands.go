package commands

import (
	"github.com/raikou/internal/ssh"
)

func RunCommand(args []string) {
	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "-l":
		ssh.Print()
		return
	}

}
