package main

import (
	"os"

	"github.com/czechbol/datbox/cmd"
)

func main() {
	err := (&cmd.Datbox{}).CobraCommand().Execute()
	if err != nil {
		os.Exit(1)
	}
}
