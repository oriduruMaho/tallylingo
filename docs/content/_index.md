---
title: "Tallylingo"
description: "A CLI tool for counting words, lines, characters, and bytes in text files"
---

# Tallylingo

![GitHub License](https://img.shields.io/github/license/oriduruMaho/tallylingo)
[![Go Report Card](https://goreportcard.com/badge/github.com/oriduruMaho/tallylingo)](https://goreportcard.com/report/github.com/oriduruMaho/tallylingo)
[![Coverage Status](https://coveralls.io/repos/github/oriduruMaho/tallylingo/badge.svg?branch=main)](https://coveralls.io/github/oriduruMaho/tallylingo?branch=main)

![Version](https://img.shields.io/badge/Version-0.4.2-blue)
[![DOI](https://zenodo.org/badge/964313902.svg)](https://doi.org/10.5281/zenodo.15320893)


Tallylingo is a fast and flexible command-line tool written in Go that helps you count:

- ✅ Number of **words**
- ✅ Number of **lines**
- ✅ Number of **characters**
- ✅ Number of **bytes**

It supports multiple files and outputs aligned results with an optional total summary.

👉 [View on GitHub](https://github.com/oriduruMaho/tallylingo)

---

## Features

- 📦 Easy CLI interface
- 🔤 Supports UTF-8 characters
- 📁 Multiple file input
- 🧾 Clean aligned output
- 🌈 Colorized summary row
- 🧪 Includes test coverage and `testdata/` support

---

## Getting Started

Install from source:

```bash
git clone https://github.com/oriduruMaho/tallylingo.git
cd tallylingo
go build -o tallylingo
