package main

import (
	"bufio"
	"log"
	"os"

	"github.com/JeffreyVdb/upperprinter/pkg/printerz"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		printerz.FPrintUpper(os.Stdout, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
