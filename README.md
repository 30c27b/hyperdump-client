# hyperdump-client

Hyperdump is a simple command line utility to save, upload and share files through the internet.

Dump files with one simple command:
```
$ dump file.go
Dumped at: https://dump.example.org/example
```

To use hyperdump-client, you also need to host a [hyperdump-server](https://github.com/30c27b/hyperdump-server) instance.

## Usage

```
Usage:
    dump [-k KEY] [-o OUTPUT] [INPUT]
    dump -g [-o OUTPUT] [INPUT]
    dump -v

Options:
    -o, --output OUTPUT     Write the result to the file given at path OUTPUT.
    -k, --key KEY           Upload the dump at the url [server]/KEY.
    -g, --get               Download a dump hosted at the given INPUT.
    -v, --version           Print the current version of hyperdump.
    -c, --config            Prompts the configuration panel.

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
```

## Installation

TODO

## Copyright

Hyperdump is distributed under the [ISC License](/LICENSE).
