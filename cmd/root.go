package cmd

import (
	"os"

	"github.com/pingidentity/pingctl/cmd/docs"
	"github.com/pingidentity/pingctl/cmd/info"
	"github.com/pingidentity/pingctl/cmd/k8s"
	"github.com/pingidentity/pingctl/cmd/license"
	"github.com/pingidentity/pingctl/cmd/lint"
	"github.com/pingidentity/pingctl/cmd/pingone"
	"github.com/pingidentity/pingctl/cmd/version"
	"github.com/spf13/cobra"
)

var (
	Verbose bool
	cfgFile string
	Profile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pingctl",
	Short: "pingctl is a CLI from Ping Identity to manage PingOne and other Ping Identity projects/tools.",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(version.VersionCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pingidentity/pingctl.yaml)")
	rootCmd.PersistentFlags().StringVarP(&Profile, "profile", "p", "", "profile to use (defaults to the first profile in the configuration file)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
