package auth

import (
	"fmt"
	"log"
	"net/url"
	"os/user"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/zalando/go-keyring"
)

const service = "Hyperdump"

// Request returns the saved token or prompts for it
func Request() (string, string) {
	u, err := user.Current()
	if err != nil {
		log.Fatal("Error: Could not get the current user:", err)
	}

	var server string
	server, err = keyring.Get(service+":server", u.Name)
	if err != nil {
		return Prompt()
	}
	server = url.QueryEscape(server)

	var token string
	token, err = keyring.Get(service+":token", u.Name)
	if err != nil {
		return Prompt()
	}

	return server, token
}

// Prompt asks the user for configuration
func Prompt() (string, string) {
	var (
		server string
		token  string
	)

	fmt.Print("Hyperdump configuration prompt\n")
	fmt.Print("[go to https://github.com/30c27b/hyperdump-client for more informations]\n")

	fmt.Print("Enter Hyperdump server: ")
	fmt.Scanln(&server)

	fmt.Print("Enter Hyperdump token: ")
	t, err := terminal.ReadPassword(1)
	if err != nil {
		log.Fatal("Error: Could not read user input:", err)
	}
	token = string(t)

	u, err := user.Current()
	if err != nil {
		log.Fatal("Error: Could not get the current user:", err)
	}

	err = keyring.Set(service+":server", u.Name, server)
	if err != nil {
		log.Fatal("Error: Could not update the keychain:", err)
	}

	err = keyring.Set(service+":token", u.Name, token)
	if err != nil {
		log.Fatal("Error: Could not update the keychain:", err)
	}

	return server, token
}
