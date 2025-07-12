package search

import (
	"bytes"
	"io"
	"os/exec"
)

func run(text string) (string, string, error) {
	cmd := exec.Command("bash", "-c", text)

	// Create pipes for stdout and stderr
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	// Read separately into buffers
	var stdoutBuf, stderrBuf bytes.Buffer
	_, err = io.Copy(&stdoutBuf, stdoutPipe)
	if err != nil {
		return "", "", err
	}
	_, err = io.Copy(&stderrBuf, stderrPipe)
	if err != nil {
		return "", "", err
	}

	err = cmd.Wait()
	if err != nil {
		return "", stderrBuf.String(), err
	}

	return stdoutBuf.String(), stderrBuf.String(), nil
}
