# GoFetch v1.0

GoFetch is a lightweight command-line tool written in Go for retrieving and displaying system information in a clean, human-friendly format.

It shows details like OS, kernel, uptime, shell, desktop environment, CPU/GPU, memory usage, local IP, and more...

## Features

- 👋 Friendly greeting with user and version
- 📋 Displays detailed system information
- 🖥️ Works on Linux (and other Unix-like systems)
- 🚀 Built with Go for fast and simple execution

## Installation

Clone and build locally:

```bash
git clone https://github.com/ujaRHR/gofetch
cd gofetch
go build -o gofetch .
```

Run:

```bash
./gofetch
```

### System-wide Installation (Linux)

To use GoFetch from anywhere in your terminal, move it to a directory in your `PATH`, like `/usr/local/bin`:

```bash
sudo mv gofetch /usr/local/bin/
```

Now you can run:

```bash
gofetch
```

from anywhere.

## Usage

```bash
gofetch            # show system info
gofetch --help     # show help
gofetch --version  # show version
```