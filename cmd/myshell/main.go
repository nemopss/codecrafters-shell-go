package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprint(os.Stdout, input[:len(input)-1]+": command not found\n")
	}
}
