package main

import (
	"fmt"
	"os"
	"os/exec"
)

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  frames test")
	fmt.Println("  frames help")
	fmt.Println("")
}

func RmRfBang() {
	exec.Command("rm", "-rf", "data").CombinedOutput()
	os.Mkdir("data", 0755)
}
