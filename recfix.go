package recutils

import (
	"fmt"
	"os/exec"
)

type RecfixOperation int

const (
	Auto RecfixOperation = iota
	Check
	Sort
)

func Recfix(filename string, operation RecfixOperation, useExternalDesc bool, force bool) error {
	error := validateFilepathDoesntExistOutsideCurrentDirectory(filename)
	if error != nil {
		return error
	} else {
		params := filename
		if useExternalDesc {
			params = "--no-external " + params
		}
		if force {
			params = "--force " + params
		}
		switch operation {
		case Auto:
			params = "--auto " + params
		case Check:
			params = "--check " + params
		case Sort:
			params = "--sort " + params
		default:
			return fmt.Errorf("invalid RecfixOperation: %d", operation)
		}
		recfixCmd := exec.Command("bash", "-c", "recfix", params)
		result, err := recfixCmd.Output()
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				fmt.Printf("recfix command failed with exit code %d: %s\n", exitError.ExitCode(), string(exitError.Stderr))
				return fmt.Errorf("recfix command failed with exit code %d", exitError.ExitCode())
			}
		}
		fmt.Printf("recfix command output: %s\n", string(result))
		return nil
	}
}
