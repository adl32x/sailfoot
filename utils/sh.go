package utils

import "os/exec"

func Execute (cmd string) (string, error) {
	out, err := exec.Command("sh","-c", cmd).Output()
	return string(out), err
}