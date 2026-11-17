package tcp

import (
	"fmt"
	"net"
	"strings"
)

// Protocol lines sent by the server.
const (
	// Defined by RFC 42TAP.
	greeting       = "OK hello proto=1"
	byeReply       = "OK bye"
	nameInUse      = "ERR 201 NAME_IN_USE"
	connectedReply = "OK connected"

	// Not in the RFC: unknown command, command before CONNECT,
	// CONNECT sent twice, or a bad name.
	unknownCommand   = "ERR 400 UNKNOWN_COMMAND"
	noConnected      = "ERR 400 NOT_CONNECTED"
	alreadyConnected = "ERR 400 ALREADY_CONNECTED"
	invalidName      = "ERR 400 INVALID_NAME"
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
