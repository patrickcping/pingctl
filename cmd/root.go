package cmd

import (
	"fmt"
	"os"

	"github.com/pingidentity/pingctl/cmd/generate"
	"github.com/pingidentity/pingctl/cmd/k8s"
	"github.com/pingidentity/pingctl/cmd/license"
	"github.com/pingidentity/pingctl/cmd/version"
	"github.com/pingidentity/pingctl/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	if err := config.Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Subcommands
	rootCmd.AddCommand(k8s.K8sCmd)
	rootCmd.AddCommand(license.LicenseCmd)
	rootCmd.AddCommand(generate.GenerateCmd)
	rootCmd.AddCommand(version.VersionCmd)

	// Global flags
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pingidentity/pingctl-config)")
	rootCmd.PersistentFlags().StringVarP(&Profile, "profile", "p", "default", "set the profile to use by it's ID (defaults to the currently selected profile)")
	viper.BindPFlag("profile", rootCmd.PersistentFlags().Lookup("activeprofile"))
}
