package port

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:     `port`,
	Summary:  `utilities regarding TCP socket listening on ports (MacOS only)`,
	Commands: []*Z.Cmd{help.Cmd, findCmd, killCmd},
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

		process, err := findProcessByPort(args[0])
		if err != nil {
			return fmt.Errorf("cannot find process for port %s; %w", args[0], err)
		}

		fmt.Printf("\033[1mPID:\033[0m %d\n", process.SystemInfo.Pid)
		fmt.Printf("\033[1mNAME:\033[0m %s\n", process.Name)

		return nil
	},
}

var killCmd = &Z.Cmd{
	Name:     `kill`,
	Summary:  `kill the process using the given port`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(caller *Z.Cmd, args ...string) error {
		if runtime.GOOS != "darwin" {
			return fmt.Errorf("this command isn't supported on this operating system")
		}

		if len(args) != 1 {
			return caller.UsageError()
		}

		process, err := findProcessByPort(args[0])
		if err != nil {
			return fmt.Errorf("cannot find process for port %s; %w", args[0], err)
		}

		err = process.SystemInfo.Kill()
		if err != nil {
			return fmt.Errorf("unable to kill process using PID %d; %w", process.SystemInfo.Pid, err)
		}

		return nil
	},
}

type Process struct {
	Name       string
	SystemInfo os.Process
}

func findProcessByPort(port string) (Process, error) {
	var process Process

	// NOTE:
	// Relying on `lsof` here as I don't know how to get this information directly from MacOS.
	// I need to check how `lsof` works internally, maybe it's not that hard.
	// Note for Linux support: we can parse the `/proc` directory to find all processes
	// listening on a TCP socket and then filter them by port.
	out, err := exec.Command("lsof", "-n", "-i", fmt.Sprint(":", port)).Output()
	if err != nil {
		// FIXME:
		// `lsof` returns 1 if nothing is found. It's not really an error.
		// We should probably dismiss the error `exit status 1`.
		return process, fmt.Errorf("unable to execute lsof command; %w", err)
	}

	pid, err := strconv.Atoi(string(regexp.MustCompile(`[0-9]+`).Find(out)))
	if err != nil {
		return process, fmt.Errorf("unable to convert PID from string to int; %w", err)
	}

	systemInfo, err := os.FindProcess(pid)
	if err != nil {
		return process, fmt.Errorf("cannot find process for PID %d; %w", pid, err)
	}

	process.SystemInfo = *systemInfo
	process.Name = string(regexp.MustCompile(`[a-z]+`).Find(out))

	return process, nil
}
