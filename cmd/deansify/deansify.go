package main

import (
	"os"

	"github.com/megacoder/deansify"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:]	{
			deansify.StripFile( arg )
		}
	} else {
		deansify.StripStdin()
	}
}
