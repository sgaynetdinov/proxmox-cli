# proxmox-cli

A lightweight CLI for Proxmox VE: manage QEMU/LXC workloads and view cluster nodes.

## Install

- Build from source:
  - `go build -o proxmox-cli ./cmd/proxmox-cli`
  - `go run ./cmd/proxmox-cli --help`

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
- `proxmox-cli version` — show Proxmox VE version and `proxmox-cli` git hash

Run `proxmox-cli --help` or any subcommand with `--help` for details.

## Development

- Tests: `make test`
- Format: `make fmt`
- Tidy/vendor deps: `make tidy`
