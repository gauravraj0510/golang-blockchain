package main

import (
	"os"
	"github.com/gauravraj0510/golang-blockchain/cli"
)



// main parses the command-line arguments and runs the appropriate command. It
// initializes a CommandLine object and calls its run() method to perform the
// blockchain operations.
func main() {
	defer os.Exit(0)
	cli := cli.CommandLine{}
	cli.Run()
}