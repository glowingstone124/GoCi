package main

import (
	"fmt"
	"os"
	"time"
)

func log(input ...string) {
	f, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	logStr := ""
	for _, line := range input {
		logStr += line
	}
	logEntry := fmt.Sprintf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), logStr)
	if _, err := f.WriteString(logEntry); err != nil {
		fmt.Println("Error writing log:", err)
	}
}
