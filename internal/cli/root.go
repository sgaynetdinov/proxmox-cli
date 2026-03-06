package cli

import (
	"context"
	"fmt"
	"os"
	"strings"

	commands_cluster "proxmox-cli/internal/cli/commands/cluster"
	commands_version "proxmox-cli/internal/cli/commands/version"
	commands_vm "proxmox-cli/internal/cli/commands/vm"
	clicontext "proxmox-cli/internal/cli/context"
	"proxmox-cli/internal/cli/utils"
	"proxmox-cli/internal/proxmox"

	"github.com/spf13/cobra"
)

const rootUsageTemplate = `{{header "USAGE"}}{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} <command> <subcommand> [flags]{{end}}{{if gt (len .Aliases) 0}}

{{header "ALIASES"}}
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

{{header "EXAMPLES"}}
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}{{$cmds := .Commands}}{{if eq (len .Groups) 0}}

{{header "AVAILABLE COMMANDS"}}{{range $cmds}}{{if (and .IsAvailableCommand (ne .Name "completion") (ne .Name "help") (ne .Name "version"))}}
  {{if or (eq $.CommandPath "proxmox-cli vm") (eq $.CommandPath "proxmox-cli cluster")}}{{rpad (aliasSet .) 30 }}{{else}}{{rpad .Name .NamePadding }}{{end}} {{.Short}}{{end}}{{end}}{{if eq .CommandPath "proxmox-cli"}}

{{header "ADDITIONAL COMMANDS"}}{{range $cmds}}{{if (or (eq .Name "completion") (eq .Name "version"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{else}}{{range $group := .Groups}}

{{header .Title}}{{range $cmds}}{{if (and (eq .GroupID $group.ID) .IsAvailableCommand (ne .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if not .AllChildCommandsHaveGroup}}

{{header "ADDITIONAL COMMANDS"}}{{range $cmds}}{{if (and (eq .GroupID "") .IsAvailableCommand (ne .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

{{header "FLAGS"}}
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

{{header "GLOBAL FLAGS"}}
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

{{header "ADDITIONAL HELP TOPICS"}}{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} <command> <subcommand> --help" for more information about a command.{{end}}
`

var rootCmd = &cobra.Command{
	Use:   "proxmox-cli",
	Short: "A CLI tool for managing Proxmox VE",
	Long:  `A command line interface for interacting with Proxmox VE API`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		apiURL, username, password, err := utils.GetCredentialsFromEnv()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting credentials: %v\n", err)
			os.Exit(1)
		}

		client, err := proxmox.NewClient(cmd.Context(), apiURL, username, password)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating Proxmox client: %v\n", err)
			os.Exit(1)
		}
		ctx := context.WithValue(cmd.Context(), clicontext.ClientKey, client)
		cmd.SetContext(ctx)
	},
}

func init() {
	cobra.AddTemplateFunc("header", func(s string) string {
		return "\x1b[1m" + s + "\x1b[0m"
	})
	cobra.AddTemplateFunc("aliasSet", func(cmd *cobra.Command) string {
		if len(cmd.Aliases) == 0 {
			return cmd.Name()
		}

		return cmd.Name() + ", " + strings.Join(cmd.Aliases, ", ")
	})
	rootCmd.SetUsageTemplate(rootUsageTemplate)

	// Cluster commands
	rootCmd.AddCommand(commands_cluster.ClusterCmd)

	// VM commands
	rootCmd.AddCommand(commands_vm.VmCmd)

	// Version commands
	rootCmd.AddCommand(commands_version.VersionCmd)
}

func Execute(version, gitCommit string) {
	commands_version.SetBuildInfo(version, gitCommit)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
