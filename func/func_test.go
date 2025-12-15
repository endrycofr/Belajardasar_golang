package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

// ---------- Test untuk average() ----------
func TestAverage(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected float64
	}{
		{"Normal case", []int{17, 20, 23, 34, 55}, 29.8},
		{"Empty slice", []int{}, 0},
		{"Single value", []int{10}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := average(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %.2f, but got %.2f", tt.expected, result)
			}
		})
	}
}

// ---------- Test untuk listNames() ----------
func TestListNames(t *testing.T) {
	names := listNames()

	if len(names) != 5 {
		t.Fatalf("Expected 5 names, but got %d", len(names))
	}

	expectedFirst := "Rizqi"
	if names[0] != expectedFirst {
		t.Errorf("Expected first name %q, but got %q", expectedFirst, names[0])
	}
}

// ---------- Test untuk PrintName() ----------
func TestPrintName(t *testing.T) {
	// Redirect stdout ke buffer sementara
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal("Failed to create pipe:", err)
	}
	os.Stdout = w

	PrintName()

	// Restore stdout & read buffer
	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)

	expected := "Nama saya Rizqi Fauzi\n"
	if buf.String() != expected {
		t.Errorf("Expected %q but got %q", expected, buf.String())
	}
}
