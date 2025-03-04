## What
[![godoc badge](http://godoc.org/github.com/dpritchett/deansify?status.png)](http://godoc.org/github.com/dpritchett/deansify)
Cross-platform utility to strip [ANSI escape codes](http://en.wikipedia.org/wiki/ANSI_escape_code#Colors) from text via `STDIN` or a named file.

## Usage

![screenshot](http://i.imgur.com/1E9Lcnt.png)

#### Streaming
`cat file_with_ansi_text | deansify`

#### By filename
`deansify file_with_ansi_text`

## Installation

#### Go users
```sh
go get -u  github.com/dpritchett/deansify
go install github.com/dpritchett/deansify/cmd/deansify
```
