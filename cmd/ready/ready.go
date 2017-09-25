package main

import "io"

// 終了コード
const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
	ExitCodeNG
)

type Ready struct {
	outStream, errStream io.Writer
}
