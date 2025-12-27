package info

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

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

func GetKernalInfo() (string, error) {
	cmd := exec.Command("uname", "-rs")
	val, err := cmd.Output()

	if err != nil {
		return "", errors.New("coudn't find any kernal info")
	}

	return string(val), nil
}
