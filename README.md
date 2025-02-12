# Raikou ⚡

**Raikou** is a lightning-fast CLI tool built with Go for managing and reconnecting SSH sessions efficiently. Designed to streamline SSH workflows, Raikou allows you to quickly access and search your SSH configuration, recall past sessions, and connect seamlessly.

## Features 🚀

- **Quick SSH Connections (WIP)** – Instantly connect to servers based on your SSH config.
- **Session Recall (WIP)** – Easily list and reconnect to previous SSH sessions.
- **Config Parsing (WIP)** – Extract and display relevant details from `~/.ssh/config`.
- **Lightweight & Fast (WIP)** – Written in Go for minimal overhead and high performance.

## Installation

### Using Go

Ensure you have Go installed, then run:

```sh
go install https://github.com/balub/Raikou.git
```

### Manual Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/balub/Raikou.git
   cd raikou
   ```
2. Build the binary:
   ```sh
   go build -o rk cmd/main.go
   ```
3. Move it to a directory in your `$PATH` (optional):
   ```sh
   mv rk /usr/local/bin/
   ```

## Usage

### List SSH Hosts

```sh
rk -l
```

Displays all available SSH hosts from `~/.ssh/config`. More commands are coming soon!

## Example SSH Config File

Raikou works with your existing `~/.ssh/config`. Here’s an example:

```ini
Host personal-server
    HostName 192.168.1.100
    User myuser
    Port 22
    IdentityFile ~/.ssh/personal_key

Host work-server
    HostName work.example.com
    User workuser
    Port 2222
    IdentityFile ~/.ssh/work_key
```

---

⚡ **Raikou – Speed up your SSH workflow!**
