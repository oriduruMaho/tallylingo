package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

func helpMessage() string {
	return fmt.Sprintf(`%s [CLI_MODE_OPTIONS] <FILEs...>
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
	humanize bool
	format   string
	help     bool
}
type options struct {
	targets  *CountingTargets
	printer  *PrintOptions
	logLevel string
}

func buildFlagSet() (*flag.FlagSet, error) {
	opts := &options{targets: &CountingTargets{}, printer: &PrintOptions{}, logLevel: "info"}
	flags := flag.NewFlagSet("tallylingo", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage()) }
	flags.StringVarP(&opts.logLevel, "log", "L", "info", "Set the log level")
	flags.StringVarP(&opts.targets.words, "words", "w", false, "Count words")
	flags.StringVarP(&opts.targets.line, "lines", "l", false, "Count lines")
	flags.StringVarP(&opts.targets.characters, "characters", "c", false, "Count characters")
	flags.StringVarP(&opts.targets.bytes, "bytes", "b", false, "Count bytes")
	flags.BoolVarP(&opts.printer.help, "help", "h", false, "Print this message")
	return flags, opts
}

func hello() string {
	return "Welcome to tallylingo!"
}

func goMain(args []string) int {
	fmt.Println(hello())
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
