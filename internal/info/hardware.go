package info

import (
	"bufio"
	"errors"
	"os"
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
