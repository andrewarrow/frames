package util

import (
	"os"
	"os/exec"
)

func RmRfBang() {
	exec.Command("rm", "-rf", "data").CombinedOutput()
	os.Mkdir("data", 0755)
}

func OpenData() {
	exec.Command("open", "data").CombinedOutput()
}
