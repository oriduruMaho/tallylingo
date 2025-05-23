package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestCountMetrics(t *testing.T) {
	text := "Hello world This is a test Another line"
	reader := strings.NewReader(text)

	lines, words, chars, bytes := countMetrics(reader)

	if lines != 1 {
		t.Errorf("Expected 1 lines, got %d", lines)
	}
	if words != 8 {
		t.Errorf("Expected 8 words, got %d", words)
	}
	if chars != 39 {
		t.Errorf("Expected 39 chars, got %d", chars)
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
	output := new(bytes.Buffer)
	_, _ = output.ReadFrom(r)

	if !strings.Contains(output.String(), "test.txt") {
		t.Error("Output missing filename")
	}
	if !strings.Contains(output.String(), "3") || !strings.Contains(output.String(), "5") {
		t.Error("Output missing expected counts")
	}
}

// func TestCLIMultipleFilesWithTotal(t *testing.T) {
// 	cmd := exec.Command("./tallylingo", "-w", "testdata/sample1.txt", "testdata/sample2.txt")
// 	var out bytes.Buffer
// 	cmd.Stdout = &out

// 	err := cmd.Run()
// 	if err != nil {
// 		t.Fatalf("Command failed: %v", err)
// 	}

// 	output := out.String()

// 	if !strings.Contains(output, "Total") {
// 		t.Error("Expected 'Total' in output for multiple files")
// 	}
// }

// func TestCLICountWords(t *testing.T) {
// 	cmd := exec.Command("./tallylingo", "-w", "testdata/sample1.txt")
// 	var out bytes.Buffer
// 	cmd.Stdout = &out
// 	cmd.Stderr = os.Stderr

// 	err := cmd.Run()
// 	if err != nil {
// 		t.Fatalf("Command failed: %v", err)
// 	}

// 	output := out.String()
// 	if !strings.Contains(output, "Words") || !strings.Contains(output, "5") {
// 		t.Errorf("Unexpected output:\n%s", output)
// 	}
// }

// func Example_tallylingo() {
// 	goMain([]string{"tallylingo"})
// 	// Output:
// 	// Welcome to tallylingo!
// }

// func TestHello(t *testing.T) {
// 	got := hello()
// 	want := "Welcome to tallylingo!"
// 	if got != want {
// 		t.Errorf("hello() = %q, want %q", got, want)
// 	}
// }

func TestHelpMessage(t *testing.T) {
	goMain([]string{"tallylingo", "-h"})
	// Output :
	// tallylingo [CLI_MODE_OPTIONS] <FILEs...>
	// CLI_MODE_OPTIONS
	// 	-w, --words        Prints the number of words in the input file
	// 	-l, --lines        Prints the number of lines in the input file
	// 	-c, --characters   Prints the number of characters in the input file
	// 	-b, --bytes        Prints the number of bytes in the input file

	// 	-h, --help        Prints this message
}
