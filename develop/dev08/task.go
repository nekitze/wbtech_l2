package main

import (
	"bufio"
	"fmt"
	ps2 "github.com/mitchellh/go-ps"
	"os"
	"strconv"
	"strings"
)

func cd(args []string) {
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Wrong arguments. Use `cd [path]`")
		return
	}

	err := os.Chdir(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func pwd(args []string) {
	if len(args) > 1 {
		fmt.Fprintln(os.Stderr, "Wrong usage. `pwd` does not accept arguments")
		return
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(currentDir)
	}
}

func echo(args []string) {
	for i := 1; i < len(args); i++ {
		fmt.Printf("%s ", args[i])
	}
	fmt.Println()
}

func kill(args []string) {
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Wrong arguments. Use `kill [pid]`")
		return
	}

	pid, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	err = process.Kill()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func ps(args []string) {
	if len(args) > 1 {
		fmt.Fprintln(os.Stderr, "Wrong usage. `ps` does not accept arguments")
		return
	}

	processes, err := ps2.Processes()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("PID:\t\tName:")
	for _, process := range processes {
		fmt.Printf("%6d\t\t%s\n", process.Pid(), process.Executable())
	}
}

func main() {
	fmt.Println("======== Welcome to Mini Shell! ========")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		request := strings.Split(scanner.Text(), " ")
		switch request[0] {
		case "cd":
			cd(request)
		case "pwd":
			pwd(request)
		case "echo":
			echo(request)
		case "kill":
			kill(request)
		case "ps":
			ps(request)
		case "q":
			fmt.Println("=============== Goodbye! ===============")
			return
		}
	}
}
