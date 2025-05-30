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

func TestGoMainWithFileInput(t *testing.T) {
	// --- 準備 ---
	origStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// CLI 引数のシミュレーション（goMain expects os.Args形式）
	args := []string{"cmd", "-w", "testdata/sample1.txt"}
	goMain(args)

	// --- 出力取得 ---
	w.Close()
	os.Stdout = origStdout
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	// --- 検証 ---
	if !strings.Contains(output, "Words") {
		t.Errorf("Expected output to contain 'Words', got: %s", output)
	}
}

func TestGoMainMultipleFilesShowsTotal(t *testing.T) {
	// 標準出力をキャプチャ
	origStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// CLI 引数をシミュレート
	args := []string{"cmd", "testdata/sample1.txt", "testdata/sample2.txt"}
	goMain(args)

	// 出力取得
	w.Close()
	os.Stdout = origStdout
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "Total") {
		t.Errorf("Expected output to include 'Total' line, got: %s", output)
	}
}

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
