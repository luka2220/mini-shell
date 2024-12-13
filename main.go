package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	Basic Command Execution: Run programs by name (e.g., ls, pwd).
	Built-in Commands: Implement a few built-in commands like cd, exit, etc.
	I/O Redirection: Support input (<) and output (>) redirection.
	Piping: Implement simple piping (|) between commands.
	Error Handling: Gracefully handle invalid commands or incorrect syntax.
*/

var (
	commands = map[string]commandOperation{
		"ls": lsCmdOp,
	}
)

type commandOperation func(string) string

func lsCmdOp(string) string {
	fmt.Printf("ls command running...\n")

	return "ls command runnning"
}

func main() {
	fmt.Print("Mini-shell\n- To executre a command enter it in the prompt\n- To quit enter 'exit' or press ctrl+c\n\n")

	fmt.Println("Testing some changes.....")

	read()
}

func read() {
	buffer := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		cmd, err := buffer.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading command: %v\n", err)
			continue
		}

		// Remove \n from cmd
		cmd = cmd[:len(cmd)-1]

		if len(cmd) > 64 {
			fmt.Println("Too long, the command has a max size of 64")
			continue
		}

		switch cmd {
		case "exit":
			os.Exit(1)
		default:
			if op, in := commands[cmd]; in {
				op(cmd)
				continue
			}
			fmt.Printf("%s is an unrecognized command...\n", cmd)
		}
	}
}
