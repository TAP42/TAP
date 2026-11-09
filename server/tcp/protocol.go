package tcp

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// Protocol lines sent by the server (RFC 42TAP).
const (
	greeting       = "OK hello proto=1"
	byeReply       = "OK BYE"
	unknownCommand = "ERR 400 UNKNOWN_COMMAND"
	nameInUse      = "ERR 201 NAME_IN_USE"
	connectedReply = "OK CONNECTED"
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
func handleCommand(conn net.Conn, command string, args string, players *registry) (bool, error) {
	switch command {
	case "QUIT":
		return true, writeLine(conn, byeReply)
	case "CONNECT":
		return false, handleConnect(conn, args, players)
	default:
		return false, writeLine(conn, unknownCommand)
	}
}

func handleConnect(conn net.Conn, args string, players *registry) error {
	err := players.add(args, conn)
	if err != nil {
		log.Printf("Error")
	}
	return writeLine(conn, connectedReply)
}
