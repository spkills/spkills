package main

import "io"

// 終了コード
const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
)

type Ready struct {
	outStream, errStream io.Writer
}
