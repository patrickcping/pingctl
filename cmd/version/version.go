package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"

	// goreleaser can also pass the specific commit if you want
	commit string = ""

	noUpdateCheck bool
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "pingctl version details",
	Long:  `The current version details of the pingctl CLI, which also checks for possible version upgrades.`,
	Run: func(cmd *cobra.Command, args []string) {

		versionString := GetVersion()
		if !noUpdateCheck {
			versionString += " (verison check here)"
		}

		fmt.Printf("pingctl version %s\n", versionString)
	},
}

func init() {
	VersionCmd.Flags().BoolVar(&noUpdateCheck, "no-update-check", false, "Don't check for CLI updates")
}

func GetVersion() string {
	versionString := version
	if commit != "" {
		versionString += fmt.Sprintf(".%s", commit)
	}
	return versionString
}
