package recutils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type Rectype struct {
	Name  string
	Value string
	Enum  []string `json:",omitempty"`
}

type RecinfResponse struct {
	Record      string
	Count       int
	Rectypedefs []Rectype
	Doc         []string
	Rectypes    []Rectype
	Key         []string
	Mandatory   []string
	Singular    []string
	Allowed     []string
	Prohibited  []string
	Unique      []string
	Auto        []string
	Sort        []string
	Comments    []string
}

func Recinf(filename string) ([]RecinfResponse, error) {
	error := validateLocalFilepath(filename)
	if error != nil {
		return []RecinfResponse{}, error
	} else {
		var stderr bytes.Buffer
		recinfRecCmd := exec.Command("recinf", filename)
		recinfRecCmd.Stderr = &stderr
		output, err := recinfRecCmd.Output()
		if err != nil {
			return []RecinfResponse{}, fmt.Errorf("failed to execute recinf command: %w", stderr.String())
		}
		reclines := strings.Split(strings.TrimSpace(string(output)), "\n")
		recinfRes := []RecinfResponse{}
		for _, line := range reclines {
			thisRecinfRes := RecinfResponse{}
			lineparts := strings.Split(line, " ")
			thisRecinfRes.Count, _ = strconv.Atoi(strings.TrimSpace(lineparts[0]))
			thisRecinfRes.Record = strings.TrimSpace(strings.Join(lineparts[1:], " "))
			recinfRes = append(recinfRes, thisRecinfRes)
		}
		for i, rec := range recinfRes {
			var stderr bytes.Buffer
			recinfDescCmd := exec.Command("recinf", "-d", "-t", rec.Record, filename)
			output, err := recinfDescCmd.Output()
			recinfDescCmd.Stderr = &stderr
			if err != nil {
				return []RecinfResponse{}, fmt.Errorf("failed to execute recinf command: %w", stderr.String())
			}
			desclines := strings.Split(strings.TrimSpace(string(output)), "\n")
			for _, line := range desclines {
				if strings.HasPrefix(strings.ToLower(line), "%type:") {
					rectypestr := strings.Split(strings.TrimSpace(line[6:]), " ")
					rectype := Rectype{
						Name:  rectypestr[0],
						Value: rectypestr[1],
					}
					if len(rectypestr) > 2 {
						rectype.Enum = rectypestr[2:]
					}
					recinfRes[i].Rectypes = append(recinfRes[i].Rectypes, rectype)
				} else if strings.HasPrefix(strings.ToLower(line), "%doc:") {
					recinfRes[i].Doc = append(recinfRes[i].Doc, strings.TrimSpace(line[5:]))
				} else if strings.HasPrefix(strings.ToLower(line), "%key:") {
					recinfRes[i].Key = strings.Split(strings.TrimSpace(line[5:]), " ")
				} else if strings.HasPrefix(strings.ToLower(line), "%mandatory:") {
					recinfRes[i].Mandatory = strings.Split(strings.TrimSpace(line[11:]), " ")
				} else if strings.HasPrefix(strings.ToLower(line), "%singular:") {
					recinfRes[i].Mandatory = strings.Split(strings.TrimSpace(line[10:]), " ")
				} else if strings.HasPrefix(strings.ToLower(line), "%allowed:") {
					recinfRes[i].Allowed = strings.Split(strings.TrimSpace(line[9:]), " ")
				} else if strings.HasPrefix(strings.ToLower(line), "%prohibited:") {
					recinfRes[i].Prohibited = strings.Split(strings.TrimSpace(line[12:]), " ")
				} else if strings.HasPrefix(strings.ToLower(line), "%unique:") {
					recinfRes[i].Unique = strings.Split(strings.TrimSpace(line[8:]), " ")
				} else if strings.HasPrefix(strings.ToLower(line), "%auto:") {
					recinfRes[i].Auto = strings.Split(strings.TrimSpace(line[6:]), " ")
				} else if strings.HasPrefix(strings.ToLower(line), "%sort:") {
					recinfRes[i].Sort = strings.Split(strings.TrimSpace(line[6:]), " ")
				} else if strings.HasPrefix(line, "#") {
					recinfRes[i].Comments = append(recinfRes[i].Comments, strings.TrimSpace(strings.TrimPrefix(line, "#")))
				}
			}
		}
		return recinfRes, nil
	}
}
