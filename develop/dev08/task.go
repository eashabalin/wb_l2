package dev08

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		input = strings.TrimSuffix(input, "\n")

		if input == "exit" {
			return
		}

		commands := strings.Split(input, "|")
		for i := 0; i < len(commands); i++ {
			strings.Trim(commands[i], " ")
		}

		for _, cmd := range commands {
			handleCommand(cmd)
		}

	}
}

func handleCommand(cmd string) {
	args := strings.Fields(cmd)
	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "Not enough arguments for cd")
			return
		}
		err := os.Chdir(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Println(dir)
	case "echo":
		fmt.Println(strings.Join(args[1:], " "))
	case "kill":
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "Not enough arguments for kill")
			return
		}
		cmd := exec.Command("kill", args[1])
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	case "ps":
		cmd := exec.Command("ps", "-e", "-o", "pid,ppid,comm")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	default:
		cmds := strings.Split(cmd, "|")
		var prevCmd *exec.Cmd
		for _, cmdStr := range cmds {
			cmdArgs := strings.Fields(strings.TrimSpace(cmdStr))
			if len(cmdArgs) == 0 {
				continue
			}
			cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
			if prevCmd != nil {
				stdin, err := prevCmd.StdoutPipe()
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					break
				}
				cmd.Stdin = stdin
			}
			if cmdStr == cmds[len(cmds)-1] {
				cmd.Stdout = os.Stdout
			}
			cmd.Stderr = os.Stderr
			err := cmd.Start()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				break
			}
			prevCmd = cmd
		}
		if prevCmd != nil {
			err := prevCmd.Wait()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
}
