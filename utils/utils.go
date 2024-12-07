package utils

import (
	"fmt"
	"os"
	"time"
)

func SaveLastAnswer(lastAnswer string) {
	file, err := os.OpenFile("lastAttempts.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		file, err = os.Create("lastAttempts.txt")
		if err != nil {
			fmt.Printf("create last error: %s\n", err.Error())
			return
		}
	}
	defer file.Close()

	_, err = file.WriteString(lastAnswer + "\n")
	if err != nil {
		fmt.Printf("save last error: %s\n", err.Error())
		return
	}
}

var running = make(map[string]time.Time)

func Start() {
	running["default"] = time.Now()
}

func Submit(value string) {
	SaveLastAnswer(value)
	fmt.Println("Answer:", value)
	fmt.Println("Runtime:", time.Since(running["default"]))
}
