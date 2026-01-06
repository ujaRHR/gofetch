package info

import (
	"bufio"
	"errors"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

type CPU struct {
	Core      int
	ModelName string
}

// type Memory struct {
// 	TotalMemory      float64
// 	UsedMemory       float64
// 	MemoryPercentage float64
// 	SwapMemory       float64
// 	UsedSwap         float64
// 	SwapPercentage   float64
// }

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

func GetMemoryInfo() (Memory, error) {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return Memory{}, errors.New("couldn't find any available meminfo")
	}
	lines := strings.SplitSeq(string(data), "\n")
	info := Memory{}

	for line := range lines {
		if strings.HasPrefix(line, "MemTotal:") {
			parts := strings.Fields(line)
			memKB, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return Memory{}, errors.New("invalid memory info: MemTotal")
			}
			memGB := memKB / 1024 / 1024
			info.TotalMemory = memGB
		}

		if strings.HasPrefix(line, "MemAvailable:") {
			parts := strings.Fields(line)
			memKB, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return Memory{}, errors.New("invalid memory info: MemAvailable")
			}
			memGB := memKB / 1024 / 1024
			info.UsedMemory = info.TotalMemory - memGB
		}

		if strings.HasPrefix(line, "SwapTotal:") {
			parts := strings.Fields(line)
			memKB, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return Memory{}, errors.New("invalid swap info: SwapTotal")
			}
			memGB := memKB / 1024 / 1024
			info.SwapMemory = memGB
		}

		if strings.HasPrefix(line, "SwapFree:") {
			parts := strings.Fields(line)
			memKB, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return Memory{}, errors.New("invalid swap info: SwapFree")
			}
			memGB := memKB / 1024 / 1024
			info.UsedSwap = info.SwapMemory - memGB
		}
	}

	if info.TotalMemory > 0 {
		info.MemoryPercentage = (info.UsedMemory / info.TotalMemory) * 100
	}

	if info.SwapMemory > 0 {
		info.SwapPercentage = (info.UsedSwap / info.SwapMemory) * 100
	}

	return info, nil
}
