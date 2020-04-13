package models

import (
	"fmt"
	"os/exec"
)

type CommandResult struct {
	Error string  `json:"err"`
	Result string `json:"result"`
}

func ExeSysCommand(cmdStr string) *CommandResult {
	cmd := exec.Command("sh", "-c", addPrefix(cmdStr))
	opBytes, err := cmd.Output()
	fmt.Println(err)
	cr := &CommandResult{
		Error:err.Error(),
		Result:string(opBytes),
	}

	return cr
}


func addPrefix(c string) string{
	return "scripts/" + c
}