package main

import (
	"flag"
	"fmt"
	"gofetch/internal/cmd"
	"gofetch/internal/info"
	"gofetch/internal/ui"
	"strings"
)

var (
	showHelp    = flag.Bool("help", false, "Show help")
	showVersion = flag.Bool("version", false, "Show version")
)

const divider = "---------------------------------"

func title(s string) string {
	return ui.Bold + ui.Green + s + ui.Reset
}

func label(s string) string {
	return title(fmt.Sprintf("%-12s", s))
}

func main() {
	flag.Parse()

	if *showHelp {
		cmd.PrintHelp()
		return
	}

	if *showVersion {
		fmt.Println(cmd.Version)
		return
	}

	username, _ := info.GetUserInfo()
	os, _ := info.GetOSRelease()
	kernel, _ := info.GetKernelInfo()
	uptime, _ := info.GetUpTime()
	shell, _ := info.GetShellInfo()
	locale, _ := info.GetLocaleInfo()

	cpu, _ := info.GetCPUInfo()
	gpu, _ := info.GetGPUInfo()
	memory, _ := info.GetMemoryInfo()
	wm := info.GetWMInfo()
	ip, _ := info.GetLocalIP()

	greeting := fmt.Sprintf(
		"👋 Howdy, @%s! Welcome to %s",
		ui.Bold+ui.Purple+username+ui.Reset,
		ui.Bold+ui.Cyan+cmd.Version+ui.Reset,
	)

	fmt.Println(divider)
	fmt.Println(greeting)
	fmt.Println(divider)
	fmt.Println()

	var b strings.Builder
	fmt.Fprintf(&b,
		"%s %s\n"+
			"%s %s\n"+
			"%s %s\n"+
			"%s %s\n"+
			"%s %s\n\n"+
			"%s %d x %s\n"+
			"%s %s\n"+
			"%s %0.2f GiB / %0.2f GiB (%0.2f%%)\n"+
			"%s %0.2f GiB / %0.2f GiB (%0.2f%%)\n"+
			"%s %s\n"+
			"%s %s\n",
		label(":: OS"), os,
		label(":: Kernel"), kernel,
		label(":: Shell"), shell,
		label(":: DE"), wm,
		label(":: Uptime"), uptime,
		label(":: CPU"), cpu.Core, cpu.ModelName,
		label(":: GPU"), gpu,
		label(":: Memory"), memory.UsedMemory, memory.TotalMemory, memory.MemoryPercentage,
		label(":: Swap"), memory.UsedSwap, memory.SwapMemory, memory.SwapPercentage,
		label(":: Local IP"), ip,
		label(":: Locale"), locale,
	)

	fmt.Print(b.String())
}
