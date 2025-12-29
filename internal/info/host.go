package info

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// OS Release Infomations
// Which stored on the /etc/os-release file
func GetOSRelease() (string, error) {
	data, err := os.ReadFile("/etc/os-release")

	if err != nil {
		return "", errors.New("couldn't find the /etc/os-release file")
	}

	lines := strings.SplitSeq(string(data), "\n")

	for line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "PRETTY_NAME") {
			val := strings.TrimPrefix(line, "PRETTY_NAME=")
			val = strings.Trim(val, `"`)
			return val, nil
		}
	}

	return "", errors.New("coudn't find any OS release info")
}

// Kernal Version Info
// Utilizing the uname command
func GetKernalInfo() (string, error) {
	cmd := exec.Command("uname", "-rs")
	val, err := cmd.Output()

	if err != nil {
		return "", errors.New("coudn't find any kernal info")
	}

	return string(val), nil
}

// Uptime Informations
// Utilizing the uptime command
func GetUpTime() (string, error) {
	cmd := exec.Command("uptime", "-p")
	out, err := cmd.Output()

	if err != nil {
		return "", errors.New("coudn't find any uptime info")
	}

	uptime := strings.TrimSpace(string(out))

	if after, ok := strings.CutPrefix(uptime, "up "); ok {
		uptime = after
	}

	return uptime, nil
}

// SHELL info from the $SHELL variable
func GetShellInfo() (string, error) {
	out := os.Getenv("SHELL")

	if out == "" {
		return "", errors.New("couldn't find SHELL variable info")
	}

	shell := filepath.Base(out)

	return shell, nil
}
