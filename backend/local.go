// Copyright (c) 2017 Gorillalabs. All rights reserved.

package backend

import (
	"io"
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

type Local struct {
	Dir    string
	Writer io.Writer
}

func (b *Local) StartProcess(cmd string, args ...string) (Waiter, io.Writer, io.Reader, io.Reader, error) {
	command := exec.Command(cmd, args...)
	command.Env = os.Environ()

	stdin, err := command.StdinPipe()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "Could not get hold of the PowerShell's stdin stream")
	}

	stdout, err := command.StdoutPipe()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "Could not get hold of the PowerShell's stdout stream")
	}

	stderr, err := command.StderrPipe()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "Could not get hold of the PowerShell's stderr stream")
	}

	err = command.Start()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "Could not spawn PowerShell process")
	}

	return command, stdin, stdout, stderr, nil
}
