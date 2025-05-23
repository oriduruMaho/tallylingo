package main

import (
	"os"
	"strings"
	"testing"
)

func TestCountMetrics(t *testing.T) {
	text := "Hello world\nThis is a test\nAnother line"
	reader := strings.NewReader(text)

	lines, words, chars, bytes := countMetrics(reader)

	if lines != 3 {
		t.Errorf("Expected 3 lines, got %d", lines)
	}
	if words != 8 {
		t.Errorf("Expected 8 words, got %d", words)
	}
	if chars != 43 {
		t.Errorf("Expected 43 chars, got %d", chars)
	}
	if bytes != len(text) {
		t.Errorf("Expected %d bytes, got %d", len(text), bytes)
	}
}

func TestPrintCounts(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// 実行
	printCounts("test.txt", 3, 5, 10, 12, &CountingTargets{
		words: true, line: true, characters: false, bytes: false,
	})

	// キャプチャ
	w.Close()
	os.Stdout = old
	output := new(strings.Builder)
	_, _ = output.ReadFrom(r)

	if !strings.Contains(output.String(), "test.txt") {
		t.Error("Output missing filename")
	}
	if !strings.Contains(output.String(), "3") || !strings.Contains(output.String(), "5") {
		t.Error("Output missing expected counts")
	}
}

func Example_tallylingo() {
	goMain([]string{"tallylingo"})
	// Output:
	// Welcome to tallylingo!
}

// func TestHello(t *testing.T) {
// 	got := hello()
// 	want := "Welcome to tallylingo!"
// 	if got != want {
// 		t.Errorf("hello() = %q, want %q", got, want)
// 	}
// }

func TestHelpMessage(t *testing.T) {
	goMain([]string{"tallylingo", "-h"})
}
