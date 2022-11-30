package greetings

import (
	"fmt"
	"os/user"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:     `greetings`,
	Summary:  `say hello to the current user`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(caller *Z.Cmd, args ...string) error {
		user, err := user.Current()
		if err != nil {
			return fmt.Errorf("unable to fetch current user; %w", err)
		}

		fmt.Printf("Greetings %s!\n", user.Username)

		return nil
	},
}
