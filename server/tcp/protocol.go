package tcp

import (
	"fmt"
	"net"
	"strings"
)

// Protocol lines sent by the server (RFC 42TAP).
const (
	greeting       = "OK hello proto=1"
	byeReply       = "OK BYE"
	unknownCommand = "ERR 400 UNKNOWN_COMMAND"
)

// parseCommand splits a line into an upper-case command and its arguments.
func parseCommand(message string) (string, string) {
	command, args, _ := strings.Cut(strings.TrimSpace(message), " ")
	return strings.ToUpper(command), args
}

// writeLine sends one protocol line to the client.
func writeLine(conn net.Conn, line string) error {
	_, err := fmt.Fprintf(conn, "%s\n", line)
	return err
}

// handleCommand answers one command and reports whether the client leaves.
func handleCommand(conn net.Conn, command string) (bool, error) {
	switch command {
	case "QUIT":
		return true, writeLine(conn, byeReply)
	default:
		return false, writeLine(conn, unknownCommand)
	}
}
