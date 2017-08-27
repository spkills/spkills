package main

import "os"

func main() {
	cli := &Ready{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
