package models

import (
	"os/exec"
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
		Result:string(opBytes),
	}
	
	return cr
}



func addPrefix(c string) string{
	return "scripts/" + c
}