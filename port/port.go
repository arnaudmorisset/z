package port

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:     `port`,
	Summary:  `utilities regarding TCP socket listening on ports (MacOS only)`,
	Commands: []*Z.Cmd{help.Cmd, findCmd},
}

var findCmd = &Z.Cmd{
	Name:     `find`,
	Summary:  `find which process use a given port`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(caller *Z.Cmd, args ...string) error {
		if runtime.GOOS != "darwin" {
			return fmt.Errorf("this command isn't supported on this operating system")
		}

		if len(args) != 1 {
			return caller.UsageError()
		}

		// NOTE:
		// Relying on `lsof` here as I don't know how to get this information directly from MacOS.
		// I need to check how `lsof` works internally, maybe it's not that hard.
		// Note for Linux support: we can parse the `/proc` directory to find all processes
		// listening on a TCP socket and then filter them by port.
		out, err := exec.Command("lsof", "-n", "-i", fmt.Sprint(":", args[0])).Output()
		if err != nil {
			// FIXME:
			// `lsof` returns 1 if nothing is found. It's not really an error.
			// We should probably dismiss the error `exit status 1`.
			return fmt.Errorf("unable to execute lsof command; %w", err)
		}

		pid := regexp.MustCompile(`[0-9]+`).Find(out)
		cmd := regexp.MustCompile(`[a-z]+`).Find(out)

		fmt.Printf("\033[1mPID:\033[0m %s\n", pid)
		fmt.Printf("\033[1mCMD:\033[0m %s\n", cmd)

		return nil
	},
}
