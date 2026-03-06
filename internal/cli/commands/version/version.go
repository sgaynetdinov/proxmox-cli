package version

import (
	"fmt"
	"strings"

	clicontext "proxmox-cli/internal/cli/context"

	"github.com/spf13/cobra"
)

var cliVersion = "dev"
var cliGitCommit = "none"

func SetBuildInfo(version, gitCommit string) {
	if version != "" {
		cliVersion = version
	}
	if gitCommit != "" {
		cliGitCommit = gitCommit
	}
}

func shortCommit(hash string) string {
	h := strings.TrimSpace(hash)
	if len(h) > 7 {
		return h[:7]
	}
	return h
}

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show Proxmox VE and proxmox-cli versions",
	Long:  `Show Proxmox VE version from API and proxmox-cli build version.`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := clicontext.GetClientFromContext(cmd)
		pveVersion, err := client.GetVersion(cmd.Context())
		if err != nil {
			return fmt.Errorf("getting Proxmox VE version: %w", err)
		}

		fmt.Printf("Proxmox VE: %s\n", pveVersion.String())
		fmt.Printf("proxmox-cli: %s\n", shortCommit(cliGitCommit))
		return nil
	},
}
