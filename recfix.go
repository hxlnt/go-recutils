package recutils

import (
	"bytes"
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
		var stderr bytes.Buffer
		recfixCmd := exec.Command("bash", "-c", "recfix", params)
		recfixCmd.Stderr = &stderr
		result, err := recfixCmd.Output()
		if err != nil {
			return fmt.Errorf("recfix command failed with exit code %d", stderr.String())
		}
		fmt.Printf("recfix command output: %s\n", string(result))
		return nil
	}
}
