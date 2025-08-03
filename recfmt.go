package recutils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func (recs RecordSet) Fmt(template string, templateIsFilename bool) (string, error) {
	recsStr := Recs2string(recs.Records)
	var tempFilePath string
	var cleanup func()
	if !templateIsFilename {
		tmpFile, err := os.CreateTemp("", "template-*.txt")
		if err != nil {
			return "", fmt.Errorf("Failed to create temp file for template: %v", err)
		}
		_, err = tmpFile.WriteString(template)
		if err != nil {
			return "", fmt.Errorf("Failed to write template to temp file: %v", err)
		}
		tmpFile.Close()
		tempFilePath = tmpFile.Name()
		cleanup = func() {
			os.Remove(tempFilePath)
		}
	} else {
		tempFilePath = template
		cleanup = func() {}
	}
	defer cleanup()
	recfmtCmd := exec.Command("recfmt", "--f", tempFilePath)
	recfmtCmd.Stdin = strings.NewReader(recsStr)
	var stderr bytes.Buffer
	recfmtCmd.Stderr = &stderr
	result, err := recfmtCmd.Output()
	if err != nil {
		return "", fmt.Errorf("Failed to execute recfmt command:\n%s", stderr.String())
	}
	return string(result), nil
}
