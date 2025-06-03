package recutils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type RecordSet struct {
	Records []Record
	Error   error
}

type Recfile struct {
	Path  string
	Error error
}

type Record struct {
	Fields []Field
}

type Field struct {
	Name  string
	Value string
}

type OptionFlags struct {
	Force           bool `json:",omitempty"`
	NoExternal      bool `json:",omitempty"`
	NoAuto          bool `json:",omitempty"`
	CaseInsensitive bool `json:",omitempty"`
	Unique          bool `json:",omitempty"`
}

var DefaultOptions OptionFlags = OptionFlags{
	Force:           false,
	NoExternal:      false,
	NoAuto:          false,
	CaseInsensitive: false,
	Unique:          false,
}

type SelectionParams struct {
	Type       string `json:",omitempty"`
	Expression string `json:",omitempty"`
	Quick      string `json:",omitempty"`
	Number     []int  `json:",omitempty"`
	Random     int    `json:",omitempty"`
	Join       string `json:",omitempty"`
}

func recs2string(records []Record) string {
	var result string
	for _, record := range records {
		for _, field := range record.Fields {
			result += fmt.Sprintf("%s: %s\n", field.Name, field.Value)
		}
		result += "\n"
	}
	return result
}

func string2recs(recStr string) []Record {
	records := strings.Split(string(recStr), "\n\n")
	var results []Record
	for _, rec := range records {
		thisRec := Record{}
		line := strings.Split(rec, "\n")
		for _, l := range line {
			if strings.TrimSpace(l) != "" {
				thisField := Field{}
				tokens := strings.Split(l, ":")
				thisField.Name = strings.TrimSpace(tokens[0])
				thisField.Value = strings.TrimSpace(strings.Join(tokens[1:], ":"))
				thisRec.Fields = append(thisRec.Fields, thisField)
			}
		}
		results = append(results, thisRec)
	}
	return results
}

func validateLocalFilepath(filename string) error {
	for _, r := range filename {
		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || strings.ContainsRune("._-/", r)) {
			return fmt.Errorf("filename contains invalid character: %q", r)
		}
	}
	currDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}
	absPath, err = filepath.EvalSymlinks(absPath)
	if err != nil {
		return fmt.Errorf("failed to resolve symlink: %w", err)
	}
	rel, err := filepath.Rel(currDir, absPath)
	if err != nil {
		return fmt.Errorf("failed to get relative path: %w", err)
	}
	if strings.HasPrefix(rel, "..") {
		return fmt.Errorf("file %s is outside the current directory", filename)
	}
	return nil
}
