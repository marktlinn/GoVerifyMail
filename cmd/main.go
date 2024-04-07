package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/marktlinn/GoVerifyMail/domain"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Scanning in data now")
	for scanner.Scan() {
		if err := domain.CheckDomain(scanner.Text()); err != nil {
			log.Printf("invalid domain: %s", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan: %s", err)
	}
}
