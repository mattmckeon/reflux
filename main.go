package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// ParseCommand processes a line of text and returns the command and any arguments.
func ParseCommand(line string) (string, string, string) {
	tokens := strings.Split(line, " ")
	var command, arg1, arg2 string
	if len(tokens) > 0 {
		command = strings.TrimSpace(tokens[0])
	}
	if len(tokens) > 1 {
		arg1 = strings.TrimSpace(tokens[1])
	}
	if len(tokens) > 2 {
		arg2 = strings.TrimSpace(tokens[2])
	}
	return command, arg1, arg2
}

// ProcessCommands reads a list of commands from the given reader, directing output the given writer.
func ProcessCommands(reader io.Reader, writer io.Writer) {
	db := NewRefluxDb()
	bufReader := bufio.NewReader(reader)
	for {
		line, err := bufReader.ReadString('\n')
		if err != nil {
			return
		}
		command, arg1, arg2 := ParseCommand(line)
		switch command {
		case "GET":
			db.DoGet(arg1, writer)
		case "SET":
			db.DoSet(arg1, arg2)
		case "DELETE":
			db.DoDelete(arg1)
		case "COUNT":
			db.DoCount(arg1, writer)
		case "BEGIN":
			db.DoBegin()
		case "COMMIT":
			db.DoCommit()
		case "ROLLBACK":
			db.DoRollback(writer)
		case "END":
			return
		}
	}
}

// The main function
func main() {
	ProcessCommands(os.Stdin, os.Stdout)
}
