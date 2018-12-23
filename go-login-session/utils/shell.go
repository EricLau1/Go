package utils

import(
	"os/exec"
)

func Shell(url string) error {

	cmd := "xdg-open"
	args := []string{url}

	return exec.Command(cmd, args...).Start()
}