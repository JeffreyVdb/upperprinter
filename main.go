package main

import (
	"bufio"
	"log"
	"os"

	"github.com/JeffreyVdb/printutils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		printutils.FPrintDouble(os.Stdout, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
