package util

import (
	"os"
	"os/exec"
)

//Clear terminal
func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
