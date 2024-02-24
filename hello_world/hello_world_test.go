package main

import (
	"io"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	expected := "Hello, World!"

	// Redirect standard output to capture the printed message
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	// Reset the standard output
	w.Close()
	os.Stdout = old

	out, _ := io.ReadAll(r)
	actual := string(out)

	if actual != expected {
		t.Errorf("Expected: %s but got: %s", expected, actual)
	}
}
