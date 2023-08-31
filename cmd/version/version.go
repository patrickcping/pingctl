package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	noUpdateCheck bool
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "pingctl version details",
	Long:  `The current version details of the pingctl CLI, which also checks for possible version upgrades.`,
	Run: func(cmd *cobra.Command, args []string) {

		versionString := GetVersion()
		if !noUpdateCheck {
			versionString += " (version check here)"
		}

		fmt.Printf("pingctl version %s\n", versionString)
	},
}

func init() {
	VersionCmd.Flags().BoolVar(&noUpdateCheck, "no-update-check", false, "Don't check for CLI updates")
}

func GetVersion() string {
	// versionString := version
	// if commit != "" {
	// 	versionString += fmt.Sprintf(".%s", commit)
	// }
	// return versionString
	return "todo"
}
