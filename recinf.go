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

type Recinfo struct {
	RecName     string
	Count       int
	Doc         []string
	Rectypedefs []Rectype
	Rectypes    []Rectype
	Key         []string
	Mandatory   []string
	Singular    []string
	Allowed     []string
	Unique      []string
	Prohibited  []string
	Auto        []string
	Sort        []string
	Comments    []string
}

func (recf Recfile) Inf() ([]Recinfo, error) {
	info := []Recinfo{}
	err := validateLocalFilepath(recf.Path)
	if err != nil {
		return info, fmt.Errorf("Filepath invalid: %s", err.Error())
	} else {
		var stderr bytes.Buffer
		recinfRecCmd := exec.Command("recinf", recf.Path)
		recinfRecCmd.Stderr = &stderr
		output, err := recinfRecCmd.Output()
		if err != nil {
			return info, fmt.Errorf("failed to execute recinf command: %s", stderr.String())
		}
		reclines := strings.Split(strings.TrimSpace(string(output)), "\n")
		recinfRes := []Recinfo{}
		for _, line := range reclines {
			thisRecinfRes := Recinfo{}
			lineparts := strings.Split(line, " ")
			thisRecinfRes.Count, _ = strconv.Atoi(strings.TrimSpace(lineparts[0]))
			thisRecinfRes.RecName = strings.TrimSpace(strings.Join(lineparts[1:], " "))
			recinfRes = append(recinfRes, thisRecinfRes)
		}
		for i, record := range recinfRes {
			var stderr bytes.Buffer
			recinfDescCmd := exec.Command("recinf", "-d", "-t", record.RecName, recf.Path)
			output, err := recinfDescCmd.Output()
			recinfDescCmd.Stderr = &stderr
			if err != nil {
				return []Recinfo{}, fmt.Errorf("failed to execute recinf command: %s", stderr.String())
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
				} else if strings.HasPrefix(strings.ToLower(line), "%typedef:") {
					rectypedefstr := strings.Split(strings.TrimSpace(line[9:]), " ")
					rectypedef := Rectype{
						Name:  rectypedefstr[0],
						Value: rectypedefstr[1],
					}
					recinfRes[i].Rectypedefs = append(recinfRes[i].Rectypedefs, rectypedef)
				} else if strings.HasPrefix(strings.ToLower(line), "%doc:") {
					recinfRes[i].Doc = append(recinfRes[i].Doc, strings.TrimSpace(line[5:]))
				} else if strings.HasPrefix(strings.ToLower(line), "%key:") {
					recinfRes[i].Key = strings.Split(strings.TrimSpace(line[5:]), " ")
				} else if strings.HasPrefix(strings.ToLower(line), "%mandatory:") {
					recinfRes[i].Mandatory = strings.Split(strings.TrimSpace(line[11:]), " ")
				} else if strings.HasPrefix(strings.ToLower(line), "%singular:") {
					recinfRes[i].Singular = strings.Split(strings.TrimSpace(line[10:]), " ")
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
