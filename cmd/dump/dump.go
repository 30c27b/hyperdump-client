package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/30c27b/hyperdump-client/internal/api"
)

// Version is the current version of Hyperump
const Version = "0.0.0"

const usage = `Hyperdump :: files sharing cli

Usage:
    dump [-k KEY] [-o OUTPUT] [INPUT]
    dump -g [-o OUTPUT] [INPUT]
    dump -v

Options:
    -o, --output OUTPUT     Write the result to the file given at path OUTPUT.
    -k, --key KEY           Upload the dump at the url [server]/KEY.
    -g, --get               Download a dump hosted at the given INPUT.
    -v, --version           Print the current version of hyperdump.

INPUT defaults to the standard input.

OUTPUT defaults to the standard output.

KEY is a 1 to 32 characters string.

Example:
    $ dump -v
    Hyperdump version 0.0.0
    $ dump -c
    Hyperdump configuration prompt
    [go to https://github.com/30c27b/hyperdump-client for more informations]
    Enter Hyperdump server: https://dump.example.org
    Enter Hyperdump token:
    Configuration successfully saved!
    $ dump -k "example" main.c
    Dumped at: https://dump.example.org/example
    $ dump -g -o output.c https://dump.example.org/example
    Dump downloaded to: output.c
`

func main() {
	flag.Usage = func() { fmt.Fprintf(os.Stderr, "%s\n", usage) }

	var (
		versionFlag bool
		getFlag     bool
		keyFlag     string
		outFlag     string
	)

	flag.BoolVar(&versionFlag, "v", false, "Print the current version of hyperdump.")
	flag.BoolVar(&versionFlag, "version", false, "Print the current version of hyperdump.")

	flag.BoolVar(&getFlag, "g", false, "Download a dump hosted at the given INPUT.")
	flag.BoolVar(&getFlag, "get", false, "Download a dump hosted at the given INPUT.")

	flag.StringVar(&keyFlag, "k", "", "Upload the dump at the url [server]/KEY.")
	flag.StringVar(&keyFlag, "key", "", "Upload the dump at the url [server]/KEY.")

	flag.StringVar(&outFlag, "o", "", "Write the result to the file given at path OUTPUT.")
	flag.StringVar(&outFlag, "output", "", "Write the result to the file given at path OUTPUT.")

	flag.Parse()

	if flag.NArg() > 1 {
		log.Fatalf("Error: too many arguments.\n\n%s\n", usage)
	}

	if versionFlag {
		fmt.Printf("Hyperdump version %s\n", Version)
		os.Exit(0)
	}

	// check for invalid arguments combinations
	switch {
	case len(keyFlag) > 0:
		log.Fatalf("Error: the -k and -g arguments cannot be used together.\n\n%s\n", usage)
	}

	var in, out io.ReadWriter = os.Stdin, os.Stdout

	if input := flag.Arg(0); input != "" && input != "-" {
		f, err := os.Open(input)
		if err != nil {
			log.Fatalf("Error: failed to open input file %q: %v\n", input, err)
		}
		defer f.Close()
		in = f
	}

	if output := outFlag; output != "" && output != "-" {
		f, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
		if err != nil {
			log.Fatalf("Error: failed to open output file %q: %v\n", output, err)
		}
		out = f
	}

	switch {
	case len(keyFlag) > 0:
		api.Pull(in, out)
		break
	default:
		api.Push(in, out, keyFlag)
	}
}
