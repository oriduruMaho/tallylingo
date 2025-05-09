# tallylingo

![GitHub License](https://img.shields.io/github/license/oriduruMaho/tallylingo)
[![Go Report Card](https://goreportcard.com/badge/github.com/oriduruMaho/tallylingo)](https://goreportcard.com/report/github.com/oriduruMaho/tallylingo)
[![Coverage Status](https://coveralls.io/repos/github/oriduruMaho/tallylingo/badge.svg?branch=main)](https://coveralls.io/github/oriduruMaho/tallylingo?branch=main)

![Version](https://img.shields.io/badge/Version-0.1.2-blue)
[![DOI](https://zenodo.org/badge/964313902.svg)](https://doi.org/10.5281/zenodo.15320893)

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
