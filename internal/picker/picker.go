package picker

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func Run(options []string) (string, error) {
	results, err := run(options, false)
	if err != nil || len(results) == 0 {
		return "", err
	}

	return results[0], nil
}

func RunFrom(parentPath string) (string, error) {
	dirs, err := getSubdirs(parentPath)
	if err != nil {
		return "", err
	}
	return Run(dirs)
}

func RunMulti(options []string) ([]string, error) {
	return run(options, true)
}

func run(options []string, multi bool) ([]string, error) {
	var args []string
	if multi {
		args = append(args, "-m")
	}

	picker := exec.Command("fzf", args...)

	pickerStdin, err := picker.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdin pipe: %w", err)
	}

	var outBuffer bytes.Buffer
	picker.Stdout = &outBuffer
	picker.Stderr = os.Stderr

	if err := picker.Start(); err != nil {
		return nil, fmt.Errorf("Failed to start picker: %w", err)
	}

	go func() {
		defer pickerStdin.Close()
		for _, option := range options {
			io.WriteString(pickerStdin, option+"\n")
		}
	}()

	if err := picker.Wait(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 130 {
			return nil, nil
		}
		return nil, fmt.Errorf("Picker exited with error: %w", err)
	}

	output := strings.TrimSpace(outBuffer.String())
	if output == "" {
		return nil, nil
	}

	return strings.Split(output, "\n"), nil
}
