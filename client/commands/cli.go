// Package commands loads the command list and runs the interactive CLI.
package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/ergochat/readline"
	"github.com/fatih/color"
)

// CommandFile is the root object of data/commands.json.
type CommandFile struct {
	Commands []Command `json:"commands"`
}

// Command describes a player command and its subcommands.
type Command struct {
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Aliases     []string  `json:"aliases"`
	Usage       string    `json:"usage"`
	Summary     string    `json:"summary"`
	Args        Args      `json:"args"`
	Examples    []string  `json:"examples"`
	Subcommands []Command `json:"subcommands"`
}

// Args holds the argument count rules of a command.
type Args struct {
	Min  int  `json:"min"`
	Max  int  `json:"max"`
	Rest bool `json:"rest"`
}

// Cli prints the banner and runs the input loop until the user quits.
func Cli() error {
	printBanner()
	fmt.Println("Type HELP for commands, QUIT to exit.")
	input, err := readline.NewFromConfig(&readline.Config{Prompt: "> "})
	if err != nil {
		return fmt.Errorf("start readline: %w", err)
	}
	defer func() {
		if err := input.Close(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
		}
	}()
	for {
		line, err := input.ReadLine()
		if err != nil || !handleLine(input, strings.TrimSpace(line)) {
			return nil
		}
	}
}

// handleLine runs one input line and returns false when the user quits.
func handleLine(input *readline.Instance, line string) bool {
	switch {
	case line == "":
	case strings.EqualFold(line, "quit"):
		return false
	case strings.EqualFold(line, "help"):
		if err := help(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
		}
	case strings.EqualFold(line, "clear"):
		input.ClearScreen()
	default:
		fmt.Printf("Unknown command %q. Type HELP.\n", line)
	}
	return true
}

func help() error {
	content, err := os.ReadFile("data/commands.json")
	if err != nil {
		return fmt.Errorf("read commands file: %w", err)
	}

	var payload CommandFile
	err = json.Unmarshal(content, &payload)
	if err != nil {
		return fmt.Errorf("parse commands file: %w", err)
	}

	bold := color.New(color.Bold).SprintFunc()
	for _, cmd := range payload.Commands {
		fmt.Printf("%s %s\n", bold(cmd.Name), cmd.Summary)
	}
	return nil
}
