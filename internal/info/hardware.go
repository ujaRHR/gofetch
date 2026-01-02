package info

import (
	"bufio"
	"errors"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

type CPU struct {
	Core      int
	ModelName string
}

// CPU Info from /proc/cpuinfo file
func GetCPUInfo() (CPU, error) {
	out, err := os.ReadFile("/proc/cpuinfo")
	totalCPU := runtime.NumCPU()

	if err != nil {
		return CPU{}, errors.New("couldn't get any /proc/cpuinfo data")
	}

	data := string(out)
	scanner := bufio.NewScanner(strings.NewReader(data))

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "model name") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) != 2 {
				continue
			}
			return CPU{
				Core:      totalCPU,
				ModelName: strings.TrimSpace(parts[1]),
			}, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return CPU{}, err
	}

	return CPU{}, errors.New("model name not found in /proc/cpuinfo")
}

func GetGPUInfo() (string, error) {
	cmd := exec.Command("lspci")
	data, err := cmd.Output()

	if err != nil {
		return "", errors.New("couldn't find any GPU info using lspci command")
	}

	lines := strings.SplitSeq(string(data), "\n")

	for line := range lines {
		if strings.Contains(strings.ToLower(line), "vga") {
			parts := strings.SplitN(line, ":", 3)
			if len(parts) == 3 {
				line = parts[2]
			}

			re := regexp.MustCompile(`\s*\(rev.*\)`)
			line = re.ReplaceAllString(line, "")

			return strings.TrimSpace(line), nil
		}
	}

	return "", errors.New("no GPU found")
}
