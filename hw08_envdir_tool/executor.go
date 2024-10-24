package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	success = 0
	failure = 1
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	//if len(cmd) == 0 {
	//	fmt.Fprintf(os.Stderr, "usage: %s /path/to/env/dir command [args...]\n", cmd[0])
	//	return failure
	//}

	com := exec.Command(cmd[0], cmd[1:]...)

	// Установить переменные окружения
	for k, v := range env {
		err := os.Unsetenv(k)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error delete environment: %v\n", err)
		}
		if !v.NeedRemove {
			err := os.Setenv(k, v.Value)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error setting environment: %v\n", err)
			}
		}
		//com.Env = append(com.Env, fmt.Sprintf("%s=%s", k, v))
	}

	com.Stderr = os.Stderr
	com.Stdout = os.Stdout
	com.Stdin = os.Stdin

	// Запустить команду
	if err := com.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return exitErr.ExitCode()
		} else {
			fmt.Fprintf(os.Stderr, "error running command: %v\n", err)
			return failure
		}
	} //com.Env = os.Environ()
	return success
}
