package recutils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func (recs RecordSet) Ins(newRecords RecordSet, params SelectionParams, options OptionFlags) RecordSet {
	response := RecordSet{
		Records: recs.Records,
		Error:   recs.Error,
	}
	recsStr := Recs2string(recs.Records)
	newRecsStr := Recs2string(newRecords.Records)
	var stderr bytes.Buffer
	paramsStr, optionsStr := parseInsArgs(params, options)
	recinsCmd := exec.Command("bash", "-c", fmt.Sprintf("echo \"%s\" | recins -r \"%s\" %s %s", recsStr, newRecsStr, paramsStr, optionsStr))
	recinsCmd.Stderr = &stderr
	result, err := recinsCmd.Output()
	if err != nil {
		response.Error = fmt.Errorf("Failed to execute recins command:\n%s", stderr.String())
	}
	response.Records = string2recs(string(result))
	return response
}

func (recf Recfile) Ins(newRecords RecordSet, params SelectionParams, options OptionFlags) Recfile {
	response := Recfile{
		Path:  recf.Path,
		Error: recf.Error,
	}
	err := validateLocalFilepath(recf.Path)
	if err != nil {
		response.Error = fmt.Errorf("Filepath invalid: %s", err.Error())
		return response
	}
	newRecsStr := Recs2string(newRecords.Records)
	var stderr bytes.Buffer
	paramsStr, optionsStr := parseInsArgs(params, options)
	recinsCmd := exec.Command("bash", "-c", fmt.Sprintf("recins -r \"%s\" %s %s %s", newRecsStr, paramsStr, optionsStr, recf.Path))
	recinsCmd.Stderr = &stderr
	err = recinsCmd.Run()
	if err != nil {
		response.Error = fmt.Errorf("Failed to execute recins command:\n%s", stderr.String())
	}
	return response
}

func parseInsArgs(params SelectionParams, options OptionFlags) (string, string) {
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
	if options.NoAuto {
		optionsStr += " --no-auto"
	}
	optionsStr += " --verbose"
	paramsStr = strings.TrimSpace(paramsStr)
	optionsStr = strings.TrimSpace(optionsStr)
	return paramsStr, optionsStr
}
