package commands

import (
	"github.com/raikou/internal/ssh"
	"github.com/raikou/internal/tui"
)

func RunCommand(args []string) {
	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "-l":
		hosts, err := ssh.ParseSSHConfigFile()
		if err != nil {
			panic(err)
		}

		tui.RenderTable(hosts)
		return
	}

}
