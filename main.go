package main

import (
	"frames/generate"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]
	RmRfBang()

	if command == "test" {
		generate.Frame()
	} else if command == "args" {
	} else if command == "help" {
		PrintHelp()
	}
}
