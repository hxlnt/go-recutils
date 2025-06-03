package recutils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type FieldAction struct {
	ActionType  string // 'Set', 'Add', 'SetAdd', 'Rename', 'Delete', 'Comment'
	ActionValue string // Blank for 'Delete' and 'Comment'
}

func (recf Recfile) Set(fields []string, action FieldAction, params SelectionParams, options OptionFlags) Recfile {
	var args []string
	response := Recfile{
		Path:  recf.Path,
		Error: recf.Error,
	}
	err := validateLocalFilepath(recf.Path)
	if err != nil {
		response.Error = fmt.Errorf("Filepath invalid: %s", err.Error())
	}
	args = append(args, "-f", strings.Join(fields, ","))
	if action.ActionType == "Set" {
		args = append(args, "-s", action.ActionValue)
	} else if action.ActionType == "Add" {
		args = append(args, "-a", action.ActionValue)
	} else if action.ActionType == "SetAdd" {
		args = append(args, "-S", action.ActionValue)
	} else if action.ActionType == "Rename" {
		args = append(args, "-r", action.ActionValue)
	} else if action.ActionType == "Delete" {
		args = append(args, "-d")
	} else if action.ActionType == "Comment" {
		args = append(args, "-c")
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
	if options.Force {
		args = append(args, "--force")
	}
	if options.NoExternal {
		args = append(args, "--no-external")
	}
	args = append(args, recf.Path)
	recsetCmd := exec.Command("recset", args...)
	var stderr bytes.Buffer
	recsetCmd.Stderr = &stderr
	err = recsetCmd.Run()
	if err != nil {
		response.Error = fmt.Errorf("Failed to execute recset command:\n%s", stderr.String())
	}
	return response
}
