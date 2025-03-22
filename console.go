package main

import (
	"os"
	"os/exec"
)

func executeShellScript(scriptPath string) {

	if !pathExist(scriptPath) {
		log("Err while executing sh: No such file or directory")
	}
	cmd := exec.Command("bash", "-l", "-c", scriptPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log("Error executing shell script:", err.Error())
	} else {
		log("Success executing shell script")
	}
}
