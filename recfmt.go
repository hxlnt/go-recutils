package recutils

import (
	"bytes"
	"fmt"
	"os/exec"
)

// type Fields struct {
// 	FieldName  string
// 	FieldValue string
// }

// type Record struct {
// 	Fields []Fields
// }

func Recfmt(records []Record, template string, templateIsFilename bool) ([]string, error) {
	var responseStr []string
	for _, record := range records {
		var fieldStr, params string
		for _, field := range record.Fields {
			fieldStr += fmt.Sprintf("%s: %s%s", field.Name, field.Value, "\n")
		}
		if templateIsFilename {
			params = "--f " + template
		} else {
			params = template
		}
		var stderr bytes.Buffer
		recfmtCmd := exec.Command("bash", "-c", fmt.Sprintf("echo \"%s\" | recfmt %s", fieldStr, params))
		recfmtCmd.Stderr = &stderr
		output, err := recfmtCmd.Output()
		if err != nil {
			fmt.Println("Error executing recfmt command:", err)
			return responseStr, fmt.Errorf("failed to execute recfmt command: %w", stderr.String())
		}
		responseStr = append(responseStr, string(output))
	}
	return responseStr, nil
}
