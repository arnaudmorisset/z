package apikey

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:     `apikey`,
	Summary:  `generate a 16 chars hexadecimal string`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(caller *Z.Cmd, args ...string) error {
		key := make([]byte, 16)
		if _, err := rand.Read(key); err != nil {
			return fmt.Errorf("unable to generate api key; %w", err)
		}

		fmt.Printf("%s\n", hex.EncodeToString(key))

		return nil
	},
}
