package recutils

import (
	"bytes"
	"fmt"
	"os/exec"
)

func (recs RecordSet) Fmt(template string, templateIsFilename bool) (string, error) {
	recsStr := Recs2string(recs.Records)
	var stderr bytes.Buffer
	var params string
	var errdetail error
	if templateIsFilename {
		params = "--f " + template
	} else {
		params = template
	}
	recfmtCmd := exec.Command("bash", "-c", fmt.Sprintf("echo \"%s\" | recfmt %s", recsStr, params))
	recfmtCmd.Stderr = &stderr
	result, err := recfmtCmd.Output()
	if err != nil {
		errdetail = fmt.Errorf("Failed to execute recfmt command:\n%s", stderr.String())
	}
	return string(result), errdetail
}
