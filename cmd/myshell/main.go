package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	valid := []string{"type", "echo", "exit"}
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
			code, err := strconv.Atoi(args[0])
			if err != nil {
				os.Exit(1)
			}
			os.Exit(code)
		case "echo":
			fmt.Fprint(os.Stdout, strings.Join(args, " ")+"\n")
		case "type":
			if slices.Contains(valid, args[0]) {
				fmt.Fprint(os.Stdout, args[0]+" is a shell builtin\n")
			} else {
				fmt.Fprint(os.Stdout, args[0]+": not found\n")
			}
		default:
			fmt.Fprint(os.Stdout, cmd+": command not found\n")
		}
	}
}
