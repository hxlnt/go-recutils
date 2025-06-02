package recutils

import (
	"fmt"
	"os/exec"
	"strings"
)

func Recdel(filename string, rectype string, expr string, q string, n []int, random int, isCaseInsensitive bool, comment bool, force bool, ignoreExternal bool) error {
	var params, options string
	error := validateFilepathDoesntExistOutsideCurrentDirectory(filename)
	if error != nil {
		return error
	}
	if rectype != "" {
		params = "-t " + rectype
	}
	if expr != "" {
		params += " -e \"" + expr + "\""
	}
	if q != "" {
		params += " -q " + q
	}
	if len(n) > 0 {
		numbers := ""
		for _, num := range n {
			if numbers != "" {
				numbers += fmt.Sprintf(",%d", num)
			} else {
				numbers += fmt.Sprintf("%d", num)
			}
		}
		params += " -n " + numbers
	}
	if random > 0 {
		params += " -r " + fmt.Sprintf("%d", random)
	}
	if isCaseInsensitive {
		options += " -i"
	}
	if comment {
		options += " -c"
	}
	if force {
		options += " --force"
	}
	if ignoreExternal {
		options += " --no-external"
	}
	options = strings.TrimSpace(options)
	params = strings.TrimSpace(params)
	recdelCmd := exec.Command("bash", "-c", fmt.Sprintf("recdel %s %s %s", options, params, filename))
	fmt.Println(recdelCmd.String())
	err := recdelCmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			fmt.Printf("recdel command failed with exit code %d: %s\n", exitError.ExitCode(), string(exitError.Stderr))
			return fmt.Errorf("recdel command failed with exit code %d", exitError.ExitCode())
		}
	}
	return nil
}
