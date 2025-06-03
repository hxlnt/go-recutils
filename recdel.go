package recutils

import (
	"bytes"
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
	var stderr bytes.Buffer
	recdelCmd := exec.Command("bash", "-c", fmt.Sprintf("recdel %s %s %s", options, params, filename))
	recdelCmd.Stderr = &stderr
	err := recdelCmd.Run()
	if err != nil {
		return fmt.Errorf("recdel command failed with exit code %s", stderr.String())
	}
	return nil
}
