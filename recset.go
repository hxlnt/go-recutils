package recutils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type FieldAction int

const (
	s FieldAction = iota
	a
	S
	r
	d
	c
)

func Recset(filename string, rectype string, expr string, q string, n []int, random int, isCaseInsensitive bool, fields []string, fieldaction FieldAction, actionvalue string, force bool, ignoreExternal bool) error {
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
	if len(fields) > 0 {
		fieldStr := strings.Join(fields, ",")
		options += fmt.Sprintf(" -f %s ", fieldStr)
	}
	if force {
		options += " --force"
	}
	if ignoreExternal {
		options += " --no-external"
	}
	switch fieldaction {
	case s:
		options += " -s " + actionvalue + " "
	case a:
		options += " -a " + actionvalue + " "
	case S:
		options += " -S " + actionvalue + " "
	case r:
		options += " -r " + actionvalue + " "
	case d:
		options += " -d "
	case c:
		options += " -c "
	default:
		return fmt.Errorf("invalid FieldAction: %d", fieldaction)
	}
	options = strings.TrimSpace(options)
	params = strings.TrimSpace(params)
	var stderr bytes.Buffer
	recsetCmd := exec.Command("bash", "-c", fmt.Sprintf("recset %s %s %s %s", options, params, filename, "--verbose"))
	recsetCmd.Stderr = &stderr
	err := recsetCmd.Run()
	if err != nil {
		return fmt.Errorf("recset failed:\n%s", stderr.String())
	}
	return nil
}
