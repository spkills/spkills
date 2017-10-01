package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun_Ready(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &Ready{outStream: outStream, errStream: errStream}
	args := strings.Split("ready", " ")

	status := cli.Create(args)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOK)
	}
}
