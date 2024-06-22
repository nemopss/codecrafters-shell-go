package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}
		cmd = strings.TrimSpace(cmd)
		switch cmd {
		case "exit 0":
			os.Exit(0)
		default:
			fmt.Fprint(os.Stdout, cmd+": command not found\n")
		}
	}
}
