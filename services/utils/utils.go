package utils

import (
    "bytes"
    "errors"
    "os/exec"
)

func GetDataModelValue(dataModelName string) (string, error) {
    // Execute the command "controller-container" with the data model name argument
    cmd := exec.Command("controller-container", dataModelName+"?")
    var out bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr
    err := cmd.Run()
    if err != nil {
        return "", errors.New("error executing command: " + err.Error() + ", " + stderr.String())
    }
    // Split the output by "=" and extract the value
    parts := bytes.Split(out.Bytes(), []byte("="))
    if len(parts) < 2 {
        return "", errors.New("nothing found for data model")
    }
    value := string(parts[1])
    return value, nil
}
