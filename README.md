# Bigfiles

A simple Go utility to find and list the largest files in a directory.

## Overview

Bigfiles scans a specified directory recursively and displays the largest files sorted by size in descending order.  
It provides a user-friendly output with the file size in human-readable format, file name, and relative path.  
If file is provided, it scans parent directory

## Installation

```bash
go install github.com/kamildemocko/bigfiles/cmd/bigfiles@latest
```

## Usage

```bash
# Scan current directory
bigfiles

# Scan specific directory
bigfiles /path/to/directory

# Change the number of files displayed
bigfiles -l 10
```

## Options

- `-l` (default: 5): Set the maximum number of files to display

## Output Example

(no color in example)

```bash
scanning files...

[ 267.85MB ] Buddha's Teachings Audiobook_ THE DHAMMAPADA - FULL, Greatest Audio Books #Audio1 (192kbit_AAC).m4a » _Audio
[ 168.00MB ] Mistrovstv__pr_ce_s_DSLR.pdf » _Programs_Skills
[ 110.22MB ] 달빛조각사1권-50권.zip » _Visual_Manga\달빛조각사
[ 91.52MB ] Cthulhu-Wars-Digital-Art-Book.pdf » _Visual_Manga\Cthulhu - Wars Digital Art Book
[ 79.60MB ] Dejiny.Presova.pdf » _Visual_Manga
```

## Dependencies

- [github.com/fatih/color](https://github.com/fatih/color) - For colorized terminal output

## License

This project is licensed under the MIT License - see the LICENSE file for details.