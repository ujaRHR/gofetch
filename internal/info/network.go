package info

import (
	"errors"
	"os/exec"
	"strings"
)

func GetLocalIP() (string, error) {
	cmd := exec.Command("hostname", "-I")
	data, err := cmd.Output()

	if err != nil {
		return "", errors.New("couldn't get any network interfaces")
	}

	ips := strings.Split(string(data), " ")

	return ips[0], nil

}
