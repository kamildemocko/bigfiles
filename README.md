# Bigfiles

A simple Go utility to find and list the largest files in a directory.

## Overview

Bigfiles scans a specified directory recursively and displays the largest files sorted by size in descending order. It provides a user-friendly output with the file size in human-readable format, file name, and relative path.

## Installation

```bash
go install github.com/kamildemocko/bigfiles/cmd/bigfiles@latest
```

## Output Example

```
[ 25.34MB ] large_file.pdf » documents\reports
[ 15.67MB ] video_clip.mp4 » media\videos
[ 8.92MB ] presentation.pptx » work\presentations
[ 3.45MB ] dataset.csv » data\videos\metadata
[ 1.23MB ] image.png » pictures
```

## Dependencies

- [github.com/fatih/color](https://github.com/fatih/color) - For colorized terminal output

## License

This project is licensed under the MIT License - see the LICENSE file for details.