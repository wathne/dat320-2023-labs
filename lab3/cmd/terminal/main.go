package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Task 7: Simple Shell
//
// This task focuses on building a simple shell that accepts
// commands that run certain OS functions or programs. For OS
// functions refer to golang's built-in OS and ioutil packages.
//
// The shell should be implemented through a command line
// application; allowing the user to execute all the functions
// specified in the task. Info such as [path] are command arguments
//
// Important: The prompt of the shell should print the current directory.
// For example, something like this:
//   /Users/meling/Dropbox/work/opsys/2020/meling-stud-labs/lab3>
//
// We suggest using a space after the > symbol.
//
// Your program should be able to at least the following functions:
// 	- exit
// 		- exit the program
// 	- cd [path]
// 		- change directory to a specified path
// 	- ls
// 		- list items and files in the current path
// 	- mkdir [path]
// 		- create a directory with the specified path
// 	- rm [path]
// 		- remove a specified file or folder
// 	- create [path]
// 		- create a file with a specified name
// 	- cat [file]
// 		- show the contents of a specified file
// 			- any file, you can use the 'hello.txt' file to check if your
// 			  implementation works
// 	- help
// 		- show a list of available commands
//
// You may also implement any number of optional functions, here are some ideas:
// 	- help [command]
// 		- give additional info on a certain command
// 	- ls [path]
// 		- make ls allow for a specified path parameter
// 	- rm -r
// 		WARNING: Be aware of where you are when you try to execute this command
// 		- recursively remove a directory
// 			- meaning that if the directory contains files, remove
// 			  all the files within the directory first, then the
// 			  directory itself
// 	- calc [expression]
// 		- Simple calculator program that can calculate a given expression
// 			- example expressions could be + - * \ pow
// 	- ipconfig
// 		- show ip interfaces
// 	- history
// 		- show command history
// 		- Alternatively implement this together with pressing up on your
// 		  keypad to load the previous command
// 		- clrhistory to clear history
// 	- tail [n]
// 		- show last n lines of a file
// 	- head [n]
// 		- show first n lines of a file
// 	- writefile [text]
// 		- write specified text to a specified file
//
// 	Or, alternatively, implement your own functionality not specified as you please
//
// Additional notes:
// 	- If you want to use colors in your terminal program you can see the package
// 		"github.com/fatih/color"
//
// 	- Helper functions may lead to cleaner code
//

// Terminal contains
type Terminal struct {
	// TODO (student): Add field(s) if necessary
}

// Execute executes a given command
func (t *Terminal) Execute(command string) {
	// TODO(wathne): Preserve whitespaces in command arguments.
	substrings := strings.Fields(command)
	length := len(substrings)
	if length == 0 {
		return
	}
	switch substrings[0] {
	case "exit":
		os.Exit(0)
	case "cd":
		if length < 2 {
			return
		}
		os.Chdir(substrings[1])
	case "ls":
		files, _ := os.ReadDir(".")
		for _, file := range files {
			fmt.Println(file.Name())
		}
	case "mkdir":
		if length < 2 {
			return
		}
		os.Mkdir(substrings[1], 0777)
	case "rm":
		if length < 2 {
			return
		}
		os.Remove(substrings[1])
	case "create":
		if length < 2 {
			return
		}
		os.Create(substrings[1])
	case "cat":
		if length < 2 {
			return
		}
		file, _ := os.Open(substrings[1])
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	case "help":
		fmt.Println("List of available commands:")
		fmt.Println("exit")
		fmt.Println("  - exit the program")
		fmt.Println("cd [path]")
		fmt.Println("  - change directory to a specified path")
		fmt.Println("ls")
		fmt.Println("  - list items and files in the current path")
		fmt.Println("mkdir [path]")
		fmt.Println("  - create a directory with the specified path")
		fmt.Println("rm [path]")
		fmt.Println("  - remove a specified file or folder")
		fmt.Println("create [path]")
		fmt.Println("  - create a file with a specified name")
		fmt.Println("cat [file]")
		fmt.Println("  - show the contents of a specified file")
		fmt.Println("help")
		fmt.Println("  - show a list of available commands")
	}
}

// This is the main function of the application.
// User input should be continuously read and checked for commands
// for all the defined operations.
// See https://golang.org/pkg/bufio/#Reader and especially the ReadLine
// function.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	terminal := Terminal{}
	wd, _ := os.Getwd()
	fmt.Print(wd, "> ")
	for scanner.Scan() {
		terminal.Execute(scanner.Text())
		wd, _ = os.Getwd()
		fmt.Print(wd, "> ")
	}
}
