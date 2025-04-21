package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading env")
		return
	}

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute()
}
