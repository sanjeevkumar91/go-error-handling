package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("liquibase", "update")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run liquibase update: %v", err)
	}

	log.Println("Liquibase migration completed successfully.")
}
