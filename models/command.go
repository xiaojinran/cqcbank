package models

import (
	"fmt"
	"os"
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

func WriteRequestBody(commandName string,requestBody []byte) {
	f, err := os.Create(addPrefix(commandName))
	if err != nil {
		fmt.Println(err)
		return
	}
	
	n2, err := f.Write(requestBody)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(n2, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func addPrefix(c string) string{
	return "scripts/" + c
}