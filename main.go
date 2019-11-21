package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/brodyzera/chorgtree"
)

var username, password, root string
var head *chorgtree.Node

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Execute input
		execInput(input)
	}
}

func execInput(input string) {
	input = strings.TrimSuffix(input, "\r\n")

	// Split input in to separate arguments
	args := strings.Split(input, " ")

	// Check for built in commands
	switch args[0] {
	case "init":
		fmt.Print("Username: ")
		fmt.Scanln(&username)

		fmt.Print("Password: ")
		fmt.Scanln(&password)

		fmt.Print("Root: ")
		fmt.Scanln(&root)

		head = chorgtree.InitTree(root, username, password)
		fmt.Println(head)
	case "print":
		b, err := json.MarshalIndent(head, "", "    ")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if len(args) > 2 && args[1] == "-o" {
			f, err := os.Create(args[2])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			defer f.Close()

			bytes, err := f.Write(b)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			fmt.Printf("wrote %d bytes\n", bytes)
		} else {
			fmt.Printf("%s\n", b)
		}
	case "load":
		if len(args) > 1 {
			data, err := ioutil.ReadFile(args[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			if head == nil {
				head = &chorgtree.Node{}
			}

			err = json.Unmarshal(data, head)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			fmt.Printf("read %d bytes\n", len(data))
		}
	case "exit":
		os.Exit(0)
	default:
		errorMessage := fmt.Sprint("command \"", args[0], "\" does not exist")
		fmt.Fprintln(os.Stderr, errorMessage)
	}
}
