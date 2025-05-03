package main

import (
	"io"
	"os"
	"os/exec"
	"time"
)

func executeShellScript(scriptPath string) {
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	log("[ShellCommandExecutor]Start executing shell command ", scriptPath, " at ", time.Now().Format("2006-01-02 15:04:05"))
	if !pathExist(scriptPath) {
		log("Err while executing sh: No such file or directory")
	}
	if err := os.MkdirAll("logs", 0755); err != nil {
		log("Cannot create logs directory")
		return
	}
	logFilePath := "logs/" + timestamp + ".log"
	f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log("Cannot open " + logFilePath)
	}
	defer f.Close()
	mw := io.MultiWriter(os.Stdout, f)
	mwErr := io.MultiWriter(os.Stderr, f)
	cmd := exec.Command("/bin/bash", scriptPath)
	stdinLog := io.TeeReader(os.Stdin, mw)
	cmd.Stdin = stdinLog
	cmd.Stdout = mw
	cmd.Stderr = mwErr
	if err := cmd.Run(); err != nil {
		log("Error executing shell script:", err.Error())
	} else {
		log("Success executing shell script")
	}
}
