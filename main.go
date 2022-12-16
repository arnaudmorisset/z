package main

import (
	"github.com/arnaudmorisset/z/apikey"
	"github.com/arnaudmorisset/z/port"
	"github.com/arnaudmorisset/z/uuid"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:      `z`,
	Summary:   `Arnaud Morisset's personal Bonzai composite command tree`,
	Version:   `v0.1.0`,
	Copyright: `Copyright 2022 Arnaud Morisset`,
	License:   `Apache-2.0`,
	Source:    `https://github.com/arnaudmorisset/z`,
	Issues:    `https://github.com/arnaudmorisset/z/issues`,
	Commands:  []*Z.Cmd{help.Cmd, apikey.Cmd, uuid.Cmd, port.Cmd},
}

func main() {
	Cmd.Run()
}
