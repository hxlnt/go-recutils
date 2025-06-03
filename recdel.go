package recutils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type DeleteStyle int

const (
	Remove DeleteStyle = iota
	Comment
)

func (recs RecordSet) Del(params SelectionParams, options OptionFlags, removeOrComment DeleteStyle) RecordSet {
	response := RecordSet{
		Records: recs.Records,
		Error:   recs.Error,
	}
	recsStr := Recs2string(recs.Records)
	var stderr bytes.Buffer
	paramsStr, optionsStr := parseDelArgs(params, options, removeOrComment)
	recdelCmd := exec.Command("bash", "-c", fmt.Sprintf("echo \"%s\" | recdel %s %s", recsStr, paramsStr, optionsStr))
	recdelCmd.Stderr = &stderr
	result, err := recdelCmd.Output()
	if err != nil {
		response.Error = fmt.Errorf("Failed to execute recdel command:\n%s", stderr.String())
	}
	response.Records = string2recs(string(result))
	return response
}

func (recf Recfile) Del(params SelectionParams, options OptionFlags, removeOrComment DeleteStyle) Recfile {
	response := Recfile{
		Path:  recf.Path,
		Error: recf.Error,
	}
	err := validateLocalFilepath(recf.Path)
	if err != nil {
		response.Error = fmt.Errorf("Filepath invalid: %s", err.Error())
	}
	var stderr bytes.Buffer
	paramsStr, optionsStr := parseDelArgs(params, options, removeOrComment)
	recdelCmd := exec.Command("bash", "-c", fmt.Sprintf("recdel %s %s %s", optionsStr, paramsStr, recf.Path))
	recdelCmd.Stderr = &stderr
	err = recdelCmd.Run()
	if err != nil {
		response.Error = fmt.Errorf("Failed to execute recdel command:\n%s", stderr.String())
	}
	return response
}

func parseDelArgs(params SelectionParams, options OptionFlags, removeOrComment DeleteStyle) (string, string) {
	var paramsStr, optionsStr string
	if params.Type != "" {
		paramsStr = "-t " + params.Type
	}
	if params.Expression != "" {
		paramsStr += " -e \"" + params.Expression + "\""
	}
	if params.Quick != "" {
		paramsStr += " -q " + params.Quick
	}
	if len(params.Number) > 0 {
		numbersSlice := make([]string, len(params.Number))
		for i, num := range params.Number {
			numbersSlice[i] = strconv.Itoa(num)
		}
		numbers := strings.Join(numbersSlice, ",")
		paramsStr += " -n " + numbers
	}
	if params.Random > 0 {
		paramsStr += " -m " + fmt.Sprintf("%d", params.Random)
	}
	if options.CaseInsensitive {
		optionsStr += " -i"
	}
	if options.Force {
		optionsStr += " --force"
	}
	if options.NoExternal {
		optionsStr += " --no-external"
	}
	if removeOrComment == Comment {
		optionsStr += " -c"
	}
	paramsStr = strings.TrimSpace(paramsStr)
	optionsStr = strings.TrimSpace(optionsStr)
	return paramsStr, optionsStr
}
