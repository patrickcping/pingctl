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
)

// versionCmd represents the version command
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "pingctl version details",
	Long:  `The current version details of the pingctl CLI, which also checks for possible version upgrades.`,
	Run: func(cmd *cobra.Command, args []string) {
		versionString := version
		if commit != "" {
			versionString += fmt.Sprintf(".%s", commit)
		}

		fmt.Printf("pingctl version %s (update logic here)\n", versionString)
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
