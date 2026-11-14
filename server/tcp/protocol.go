package tcp

import (
	"fmt"
	"net"
	"strings"
)

// Protocol lines sent by the server (RFC 42TAP).
const (
	greeting       = "OK hello proto=1"
	byeReply       = "OK bye"
	unknownCommand = "ERR 400 UNKNOWN_COMMAND"
	nameInUse      = "ERR 201 NAME_IN_USE"
	connectedReply = "OK connected"
	noConnected    = "ERR 400 NOT_CONNECTED"

	// Not in the RFC: CONNECT sent twice, or with a bad name.
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

// handleCommand answers one command and reports whether the client leaves.
func handleCommand(
	sess *session,
	command string,
	args string,
	players *registry,
) (bool, error) {
	switch command {
	case "QUIT":
		return true, writeLine(sess.conn, byeReply)
	case "CONNECT":
		return false, handleConnect(sess, args, players)
	}
	if sess.player == nil {
		return false, writeLine(sess.conn, noConnected)
	}
	return false, writeLine(sess.conn, unknownCommand)
}

// handleConnect gives the session a player name, once, if the name is free.
func handleConnect(sess *session, name string, players *registry) error {
	if sess.player != nil {
		return writeLine(sess.conn, alreadyConnected)
	}
	if name == "" || strings.ContainsAny(name, " \t") {
		return writeLine(sess.conn, invalidName)
	}
	player, err := players.add(name, sess.conn)
	if err != nil {
		return writeLine(sess.conn, nameInUse)
	}
	sess.player = player
	return writeLine(sess.conn, connectedReply)
}
