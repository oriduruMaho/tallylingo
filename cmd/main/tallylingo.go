package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	flag "github.com/spf13/pflag"
)

func helpMessage() string {
	return (`tallylingo [CLI_MODE_OPTIONS] <FILEs...>
CLI_MODE_OPTIONS
  -w, --words        Prints the number of words in the input file
  -l, --lines        Prints the number of lines in the input file
  -c, --characters   Prints the number of characters in the input file
  -b, --bytes        Prints the number of bytes in the input file

  -h, --help        Prints this message`)
}

type CountingTargets struct {
	words      bool
	line       bool
	characters bool
	bytes      bool
}

type PrintOptions struct {
	// humanize bool
	// format   string
	help bool
}

type options struct {
	targets  *CountingTargets
	printer  *PrintOptions
	logLevel string
}

func buildFlagSet() (*flag.FlagSet, *options) {
	opts := &options{targets: &CountingTargets{}, printer: &PrintOptions{}, logLevel: "info"}
	flags := flag.NewFlagSet("tallylingo", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage()) }
	flags.StringVarP(&opts.logLevel, "log", "L", "info", "Set the log level")
	flags.BoolVarP(&opts.targets.words, "words", "w", false, "Count words")
	flags.BoolVarP(&opts.targets.line, "lines", "l", false, "Count lines")
	flags.BoolVarP(&opts.targets.characters, "characters", "c", false, "Count characters")
	flags.BoolVarP(&opts.targets.bytes, "bytes", "b", false, "Count bytes")
	flags.BoolVarP(&opts.printer.help, "help", "h", false, "Print this message")
	return flags, opts
}

func countMetrics(r io.Reader) (lines int, words int, chars int, bytes int) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		lines++
		chars += len([]rune(line))
		bytes += len(line)

		wordScanner := bufio.NewScanner(strings.NewReader(line))
		wordScanner.Split(bufio.ScanWords)
		for wordScanner.Scan() {
			words++
		}
	}

	return
}

func parseAndValidateFlags(flags *flag.FlagSet, opts *options, args []string) error {
	err := flags.Parse(args[1:])
	if err != nil {
		return fmt.Errorf("error parsing flags: %w", err)
	}

	if !opts.targets.words && !opts.targets.line && !opts.targets.characters && !opts.targets.bytes {
		opts.targets.words = true
		opts.targets.line = true
		opts.targets.characters = true
		opts.targets.bytes = true
	}
	return nil
}

func printCountsHeader(opts *CountingTargets) {
	fmt.Printf("%-15s", "File")
	if opts.line {
		fmt.Printf("%10s", "Lines")
	}
	if opts.words {
		fmt.Printf("%10s", "Words")
	}
	if opts.characters {
		fmt.Printf("%10s", "Chars")
	}
	if opts.bytes {
		fmt.Printf("%10s", "Bytes")
	}
	fmt.Println()
}

func printCounts(filename string, lines, words, chars, bytes int, opts *CountingTargets) {
	isTotal := filename == "Total"

	colorStart := ""
	colorEnd := ""

	if isTotal {
		colorStart = "\033[36m" // シアン
		colorEnd = "\033[0m"
	}

	fmt.Printf("%s", colorStart)
	fmt.Printf("%-15s", filename) // 左詰め15文字分のファイル名欄

	if opts.line {
		fmt.Printf("%10d", lines)
	}
	if opts.words {
		fmt.Printf("%10d", words)
	}
	if opts.characters {
		fmt.Printf("%10d", chars)
	}
	if opts.bytes {
		fmt.Printf("%10d", bytes)
	}
	fmt.Printf("%s\n", colorEnd)
}

type countTotals struct {
	lines int
	words int
	chars int
	bytes int
}

func processFiles(files []string, targets *CountingTargets) countTotals {
	var total countTotals

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open file %s: %v\n", file, err)
			continue
		}
		defer f.Close()

		lines, words, chars, bytes := countMetrics(f)
		printCounts(file, lines, words, chars, bytes, targets)

		total.lines += lines
		total.words += words
		total.chars += chars
		total.bytes += bytes
	}
	return total
}

func goMain(args []string) int {
	flags, opts := buildFlagSet()
	if err := parseAndValidateFlags(flags, opts, args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	// ヘルプフラグが指定された場合は、ここでヘルプメッセージを表示して終了
	if opts.printer.help {
		flags.Usage()
		return 0
	}

	files := flags.Args()
	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "No input files specified.")
		return 1
	}

	printCountsHeader(opts.targets)

	total := processFiles(files, opts.targets)

	if len(files) > 1 {
		printCounts("Total", total.lines, total.words, total.chars, total.bytes, opts.targets)
	}

	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
