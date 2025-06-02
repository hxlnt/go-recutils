package recutils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func Recins(filename string, rectype string, expr string, q string, n []int, random int, isCaseInsensitive bool, record Record, force bool, ignoreExternal bool, ignoreAuto bool) error {
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
	if len(record.Fields) > 0 {
		for _, field := range record.Fields {
			options += fmt.Sprintf(" -f %s -v \"%s\"", field.FieldName, field.FieldValue)
		}
	}
	if force {
		options += " --force"
	}
	if ignoreExternal {
		options += " --no-external"
	}
	if ignoreAuto {
		options += " --no-auto"
	}
	options = strings.TrimSpace(options)
	params = strings.TrimSpace(params)
	var stderr bytes.Buffer
	recinsCmd := exec.Command("bash", "-c", fmt.Sprintf("recins %s %s %s %s", options, params, filename, "--verbose"))
	recinsCmd.Stderr = &stderr
	err := recinsCmd.Run()
	if err != nil {
		return fmt.Errorf("recins failed:\n%s", stderr.String())
	}
	return nil
}
