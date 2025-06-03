package recutils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func (recf Recfile) Sel(sortBy []string, groupBy []string, params SelectionParams, options OptionFlags) RecordSet {
	var args []string
	response := RecordSet{}
	if sortBy != nil {
		args = append(args, "-S", strings.Join(sortBy, ","))
	}
	if groupBy != nil {
		args = append(args, "-G", strings.Join(groupBy, ","))
	}
	if params.Type != "" {
		args = append(args, "-t", params.Type)
	}
	if params.Expression != "" {
		args = append(args, "-e", params.Expression)
	}
	if params.Quick != "" {
		args = append(args, "-q", params.Quick)
	}
	if params.Join != "" {
		args = append(args, "-j", params.Join)
	}
	if len(params.Number) > 0 {
		numbersSlice := make([]string, len(params.Number))
		for i, num := range params.Number {
			numbersSlice[i] = strconv.Itoa(num)
		}
		args = append(args, "-n", strings.Join(numbersSlice, ","))
	}
	if params.Random > 0 {
		args = append(args, "-m", fmt.Sprintf("%d", params.Random))
	}
	if options.CaseInsensitive {
		args = append(args, "-i")
	}
	args = append(args, recf.Path)
	recselCmd := exec.Command("recsel", args...)
	var stderr bytes.Buffer
	recselCmd.Stderr = &stderr
	result, err := recselCmd.Output()
	if err != nil {
		response.Error = fmt.Errorf("Failed to execute recsel command:\n%s", stderr.String())
	}
	response.Records = string2recs(string(result))
	return response
}

func GroupBy(strings []string) []string {
	// if strings == nil {
	// 	return []string{}
	// }
	return strings
}

func SortBy(strings []string) []string {
	// if strings == nil {
	// 	return []string{}
	// }
	return strings
}
