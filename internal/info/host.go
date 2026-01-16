package info

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// User Info via whoami command
func GetUserInfo() (string, error) {
	cmd := exec.Command("whoami")
	data, err := cmd.Output()

	if err != nil {
		return "", errors.New("couldn't find username via `whoami`")
	}

	username := strings.TrimSpace(string(data))

	return username, nil
}

// OS Release Infomations
// Which stored on the /etc/os-release file
func GetOSRelease() (string, error) {
	data, err := os.ReadFile("/etc/os-release")

	if err != nil {
		return "", errors.New("couldn't find the /etc/os-release file")
	}

	lines := strings.SplitSeq(string(data), "\n")
	var prettyName, releaseType string

	for line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "PRETTY_NAME") {
			val := strings.TrimPrefix(line, "PRETTY_NAME=")
			prettyName = strings.Trim(val, `"`)
		}

		if strings.HasPrefix(line, "RELEASE_TYPE") {
			val := strings.TrimPrefix(line, "RELEASE_TYPE=")
			releaseType = strings.Trim(val, `"`)
		}
	}

	return fmt.Sprintf("%s / %s", prettyName, releaseType), nil
}

// Kernal Version Info
// Utilizing the uname command
func GetKernelInfo() (string, error) {
	cmd := exec.Command("uname", "-rs")
	data, err := cmd.Output()

	if err != nil {
		return "", errors.New("coudn't find any kernel info")
	}

	kernel := strings.TrimSpace(string(data))
	return kernel, nil
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
		return "", errors.New("couldn't find $SHELL info")
	}

	shell := filepath.Base(out)

	return shell, nil
}

// Locale info from the $LANG env variable
func GetLocaleInfo() (string, error) {
	locale := os.Getenv("LANG")

	if locale == "" {
		return "", errors.New("couldn't find $LANG info")
	}

	return locale, nil
}
