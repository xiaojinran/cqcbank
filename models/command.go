package models

import (
	"os/exec"
	"strings"
)

type CommandResult struct {
	Error string  `json:"err"`
	Result interface{} `json:"result"`
}

func ExeSysCommand(cmdStr string) *CommandResult {
	cmd := exec.Command("sh", "-c", addPrefix(cmdStr))
	opBytes, err := cmd.Output()
	if err != nil {
		cr := &CommandResult{
			Error: err.Error(),
			Result: "",
		}
		return cr
	}
	cr := &CommandResult{
		Error: "no errors",
		Result:strings.Replace(string(opBytes), "\n", "", -1),
	}
	
	return cr
}



func addPrefix(c string) string{
	return "scripts/" + c
}