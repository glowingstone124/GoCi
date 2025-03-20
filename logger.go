package main

import (
	"fmt"
	"os"
	"time"
)

func log[T any](input ...T) {
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer f.Close()

	logStr := ""
	for _, item := range input {
		logStr += fmt.Sprint(item)
	}

	logEntry := fmt.Sprintf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), logStr)

	if _, err := f.WriteString(logEntry); err != nil {
		fmt.Println("Error writing log:", err)
	}
}
