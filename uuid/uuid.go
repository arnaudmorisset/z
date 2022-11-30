package uuid

import (
	"fmt"

	"github.com/google/uuid"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:     `uuid`,
	Summary:  `generate a random UUID v4`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(caller *Z.Cmd, args ...string) error {
		fmt.Println(uuid.New().String())

		return nil
	},
}
