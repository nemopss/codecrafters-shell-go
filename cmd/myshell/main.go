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
		mainCmd := strings.Split(cmd, " ")[0]
		args := strings.Split(cmd, " ")[1:]
		switch mainCmd {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Fprint(os.Stdout, strings.Join(args, " ")+"\n")
		default:
			fmt.Fprint(os.Stdout, cmd+": command not found\n")
		}
	}
}
