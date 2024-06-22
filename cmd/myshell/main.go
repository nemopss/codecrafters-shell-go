package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func main() {
	valid := []string{"type", "echo", "exit", "pwd", "cd"}
	paths := strings.Split(os.Getenv("PATH"), ":")
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}

		input = strings.TrimRight(input, "\n")
		tokenizedInput := strings.Split(input, " ")
		cmd := tokenizedInput[0]
		params := tokenizedInput[1:]
		switch cmd {
		case "exit":
			doExit(params)
		case "echo":
			doEcho(params)
		case "type":
			doType(params, valid, paths)
		case "pwd":
			doPwd()
		case "cd":
			doCd(params)
		default:
			doCmd(cmd, params)
		}
	}

}

func doExit(params []string) {
	code, err := strconv.Atoi(params[0])
	if err != nil || len(params) != 1 {
		os.Exit(1)
	}
	os.Exit(code)
}

func doEcho(params []string) {
	fmt.Fprint(os.Stdout, strings.Join(params, " ")+"\n")
}

func doType(params, valid, paths []string) {
	if len(params) != 1 {
		os.Exit(1)
	}
	if slices.Contains(valid, params[0]) {
		fmt.Fprint(os.Stdout, params[0]+" is a shell builtin\n")
	} else {
		for _, path := range paths {
			fp := filepath.Join(path, params[0])
			if _, err := os.Stat(fp); err == nil {
				fmt.Fprint(os.Stdout, params[0]+" is "+fp+"\n")
				return
			}
		}
		fmt.Fprint(os.Stdout, params[0]+": not found\n")

	}
}

func doCmd(cmd string, params []string) {
	command := exec.Command(cmd, params...)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout

	err := command.Run()
	if err != nil {
		fmt.Fprint(os.Stdout, cmd+": command not found\n")
	}
}

func doPwd() {
	wd, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}
	fmt.Fprint(os.Stdout, wd+"\n")
}

func doCd(params []string) {
	if len(params) != 1 {
		os.Exit(1)
	}
	err := os.Chdir(params[0])
	if err != nil {
		fmt.Fprint(os.Stdout, "cd: "+params[0]+": No such file or directory\n")
	}
}
