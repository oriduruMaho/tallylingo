# tallylingo

![GitHub License](https://img.shields.io/github/license/oriduruMaho/tallylingo)

This is a CLI tool that allows you to easily obtain text statistics such as the number of characters, words, and lines.

## Description
The program counts the number of characters, lines, words, and bytes in the specified text file.
It is developed in Go. It supports input from several file formats.

## Usage
```
tallylingo version
tallylingo [CLI_MODE_OPTIONS] <FILEs...>
CLI_MODE_OPTIONS
  -w, --words        Prints the number of words in the input file
  -l, --lines        Prints the number of lines in the input file
  -c, --characters   Prints the number of characters in the input file
  -b, --bytes        Prints the number of bytes in the input file

  -h, --help        Prints this message
```
