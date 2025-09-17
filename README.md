# proxmox-cli

A lightweight command-line tool to interact with Proxmox VE: list and control virtual machines, and view cluster nodes.

## Install

- Build from source:
  - `go build -o proxmox-cli .`
  - `./proxmox-cli --help`

Note: Package managers and release binaries may be added later.

## Configuration

Set the following environment variables (loaded at startup):

- `PM_API_URL` — Proxmox API URL (e.g. `https://pve.example:8006/api2/json`)
- `PM_USER` — username with realm (e.g. `root@pam`)
- `PM_PASS` — password

`.env` files are supported and auto-loaded. Example `.env`:

```
PM_API_URL=https://pve.example:8006/api2/json
PM_USER=root@pam
PM_PASS=your_password
```

## Usage

After configuring environment variables (or `.env`), run commands like:

- `proxmox-cli vm ps -a` — list all VMs
- `proxmox-cli vm start 101 102` — start VMs 101 and 102
- `proxmox-cli vm shutdown 203 --force` — force stop VM 203
- `proxmox-cli cluster list` — list cluster nodes

### Commands (summary)

- `vm ps [-a]` — list VMs; aliases: `list`, `ls`
- `vm start <VM_ID...>` — start one or more VMs
- `vm stop <VM_ID...>` — hard stop one or more VMs
- `vm shutdown <VM_ID...> [--force|-f]` — graceful shutdown; `--force` uses hard stop
- `vm reboot <VM_ID...> [--force|-f]` — graceful reboot; `--force` uses hard reset
- `vm reset <VM_ID...>` — hard reset (not supported for LXC)
- `cluster list` — list cluster nodes; aliases: `ls`, `ps`

Run `proxmox-cli --help` or any subcommand with `--help` for details.

## Development

- Run: `make run`
- Tests: `make test`
- Tidy/vendor deps: `make tidy`
