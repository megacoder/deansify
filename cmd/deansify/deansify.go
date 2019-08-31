package main

import (
	"os"

	"github.com/megacoder/deansify"
)

func main() {
	if len(os.Args) > 1 {
		fileName := os.Args[1]
		deansify.StripFile(fileName)
	} else {
		deansify.StripStdin()
	}
}
