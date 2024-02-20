package cmd

import (
	"io"
	"os"
)

type Configuration struct {
	Out    io.Writer
	OutErr io.Writer
}

var Config = Configuration{
	Out:    os.Stdout,
	OutErr: os.Stderr,
}
