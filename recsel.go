package recutils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func Recsel(filename string, rectype string, expr string, q string, n []int, random int, isCaseInsensitive bool, joinfield string, sortbyfields []string, groupbyfields []string, removeDuplicates bool, force bool, ignoreExternal bool) ([]Record, error) {
	var params, options string
	var results []Record
	error := validateLocalFilepath(filename)
	if error != nil {
		return results, error
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
	if joinfield != "" {
		params += " -j " + joinfield
	}
	if len(sortbyfields) > 0 {
		sortFields := strings.Join(sortbyfields, ",")
		params += " -S " + sortFields
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
		params += " -m " + fmt.Sprintf("%d", random)
	}
	if isCaseInsensitive {
		options += " -i"
	}
	if len(groupbyfields) > 0 {
		fieldStr := strings.Join(groupbyfields, ",")
		options += fmt.Sprintf(" -G %s", fieldStr)
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
	recselCmd := exec.Command("bash", "-c", fmt.Sprintf("recsel %s %s %s", options, params, filename))
	fmt.Println(recselCmd.String())
	recselCmd.Stderr = &stderr
	output, err := recselCmd.Output()
	if err != nil {
		return results, fmt.Errorf("recset failed:\n%s", stderr.String())
	}
	results = string2recs(string(output))
	// records := strings.Split(string(output), "\n\n")
	// for _, rec := range records {
	// 	thisRec := Record{}
	// 	line := strings.Split(rec, "\n")
	// 	for _, l := range line {
	// 		if strings.TrimSpace(l) != "" {
	// 			thisField := Field{}
	// 			tokens := strings.Split(l, ":")
	// 			thisField.Name = strings.TrimSpace(tokens[0])
	// 			thisField.Value = strings.TrimSpace(strings.Join(tokens[1:], ":"))
	// 			thisRec.Fields = append(thisRec.Fields, thisField)
	// 		}
	// 	}
	// 	results = append(results, thisRec)
	// }
	return results, nil
}
