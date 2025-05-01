package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestMainIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	cmd := exec.Command("go", "run", ".", "--year", "2022")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run main: %v\nOutput: %s", err, output)
	}

	expected := "Go-ldenweek is 2022/4/29 ~ 2022/5/8"
	if !strings.Contains(string(output), expected) {
		t.Errorf("Output does not contain expected string. Got: %s, Want: %s", output, expected)
	}
}
