package recutils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type FixAction int

const (
	Check FixAction = iota
	Auto
	Sort
)

func (recs RecordSet) Fix(action FixAction, options OptionFlags) RecordSet {
	response := RecordSet{
		Records: recs.Records,
		Error:   recs.Error,
	}
	recsStr := Recs2string(recs.Records)
	var stderr bytes.Buffer
	optionsStr := parseFixArgs(action, options)
	recfixCmd := exec.Command("bash", "-c", fmt.Sprintf("echo \"%s\" | recfix %s", recsStr, optionsStr))
	recfixCmd.Stderr = &stderr
	result, err := recfixCmd.Output()
	if err != nil {
		response.Error = fmt.Errorf("Recfix found validation errors:\n%s", stderr.String())
	}
	response.Records = string2recs(string(result))
	return response
}

func (recf Recfile) Fix(action FixAction, options OptionFlags) Recfile {
	response := Recfile{
		Path:  recf.Path,
		Error: recf.Error,
	}
	err := validateLocalFilepath(recf.Path)
	if err != nil {
		response.Error = fmt.Errorf("Filepath invalid: %s", err.Error())
		return response
	}
	var stderr bytes.Buffer
	optionsStr := parseFixArgs(action, options)
	recfixCmd := exec.Command("bash", "-c", fmt.Sprintf("recfix %s %s", optionsStr, recf.Path))
	recfixCmd.Stderr = &stderr
	err = recfixCmd.Run()
	if err != nil {
		response.Error = fmt.Errorf("Recfix found validation errors:\n%s", stderr.String())
	}
	return response
}

func parseFixArgs(action FixAction, options OptionFlags) string {
	var optionsStr string
	if options.Force {
		optionsStr += " --force"
	}
	if options.NoExternal {
		optionsStr += " --no-external"
	}
	switch action {
	case Auto:
		optionsStr += " --auto"
	case Sort:
		optionsStr += " --sort"
	default:
		optionsStr += " --check"
	}
	return strings.TrimSpace(optionsStr)
}
