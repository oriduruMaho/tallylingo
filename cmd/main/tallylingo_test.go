package main

import (
	"bytes"
	"os"
	"strings"
	"testing"

	flag "github.com/spf13/pflag"
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

func TestPrintCountsHeader(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printCountsHeader(&CountingTargets{
		line: true, words: true, characters: true, bytes: true,
	})

	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	for _, label := range []string{"Lines", "Words", "Chars", "Bytes"} {
		if !strings.Contains(output, label) {
			t.Errorf("Header missing: %s", label)
		}
	}
}

func TestPrintCountsTotalColor(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printCounts("Total", 3, 5, 10, 12, &CountingTargets{
		words: true, line: true, characters: false, bytes: false,
	})

	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "\033[36m") || !strings.Contains(output, "\033[0m") {
		t.Error("Expected ANSI color codes for 'Total' row")
	}
}

// func TestGoMainWithMissingFile(t *testing.T) {
// 	oldStdout := os.Stdout
// 	oldStderr := os.Stderr

// 	rOut, wOut, _ := os.Pipe()
// 	rErr, wErr, _ := os.Pipe()

// 	os.Stdout = wOut
// 	os.Stderr = wErr

// 	// 実行
// 	goMain([]string{"cmd", "-w", "nonexistent.txt"})

// 	// 復元とクローズ
// 	wOut.Close()
// 	wErr.Close()
// 	os.Stdout = oldStdout
// 	os.Stderr = oldStderr

// 	// 出力を取得
// 	var outBuf, errBuf bytes.Buffer
// 	_, _ = outBuf.ReadFrom(rOut)
// 	_, _ = errBuf.ReadFrom(rErr)

// 	// 検証
// 	if !strings.Contains(errBuf.String(), "File not found") {
// 		t.Errorf("Expected error message in stderr, got: %s", errBuf.String())
// 	}
// }

func TestParseAndValidateFlagsSetsAllDefault(t *testing.T) {
	opts := &options{targets: &CountingTargets{}, printer: &PrintOptions{}}
	flags := flag.NewFlagSet("test", flag.ContinueOnError)
	flags.BoolVarP(&opts.targets.words, "words", "w", false, "Count words")
	flags.BoolVarP(&opts.targets.line, "lines", "l", false, "Count lines")
	flags.BoolVarP(&opts.targets.characters, "characters", "c", false, "Count characters")
	flags.BoolVarP(&opts.targets.bytes, "bytes", "b", false, "Count bytes")
	flags.BoolVarP(&opts.printer.help, "help", "h", false, "Print help")

	err := parseAndValidateFlags(flags, opts, []string{"cmd"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !(opts.targets.words && opts.targets.line && opts.targets.characters && opts.targets.bytes) {
		t.Error("Expected all count flags to be set to true by default")
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
