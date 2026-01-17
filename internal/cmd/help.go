package cmd

import "fmt"

func PrintHelp() {
	fmt.Print(`
GoFetch is a lightweight Go-based fetch utility for retrieving and displaying system information.

Usage:
  gofetch
  gofetch [options]

Options:
  --help        Show this help message
  --version     Show version information
`)
}
