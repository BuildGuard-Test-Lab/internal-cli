package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: icli <command> [args]")
		fmt.Println("Commands: deploy, db, logs, status")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "deploy":
		fmt.Println("Deploying service...")
	case "status":
		fmt.Println("All services healthy")
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}
