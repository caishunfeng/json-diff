package handler

import (
	"bytes"
	"encoding/json"
	"os"
)

func prettyfy(jstr string) (string, error) {
	b := []byte(jstr)
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.String(), err
}

func writeJSONStringToFile(filename, jsonString string) error {
	fn, err := os.Create(filename)
	defer fn.Close()
	if err != nil {
		return err
	}
	prettyJSON, err := prettyfy(jsonString)
	if err != nil {
		return err
	}
	_, err = fn.WriteString(prettyJSON)
	return err
}
