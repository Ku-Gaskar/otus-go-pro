package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRunCmd(t *testing.T) {
	e, err := ReadDir("./testdata/env")
	require.NoError(t, err)

	code := RunCmd([]string{`C:\Program Files\Git\bin\bash.exe`, "-c", "echo arg=1"}, e)
	require.Equal(t, 0, code, "Exit code should be 0")
}
