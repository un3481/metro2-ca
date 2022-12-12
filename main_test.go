
package main

import (
	"bytes"
	"os"
	"testing"
)

func TestUnknown(t *testing.T) {
	_, err := executeCommand(rootCmd, "unknown")
	if err == nil {
		t.Errorf("don't support unknown")
	}
}

func TestWeb(t *testing.T) {
	_, err := executeCommand(rootCmd, "web", "--test=true")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestWebWithInvalidPort(t *testing.T) {
	_, err := executeCommand(rootCmd, "web", "--port", "unknown", "--test=false")
	if err == nil {
		t.Errorf("Starting web server with unknown port")
	}
}
