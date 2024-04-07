package main

import (
	"bufio"
	"log"
	"os"

	"github.com/marktlinn/GoVerifyMail/domain"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		domain.CheckDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan: %s", err)
	}
}
