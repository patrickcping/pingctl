package config

// import (
// 	"github.com/pingidentity/pingctl/internal/pingone"
// 	"github.com/spf13/cobra"
// )

// var (
// 	pingoneConfig pingone.ProfilePingOneCmdAuthConfig
// )

// // configCmd represents the config command
// var ConfigCmd = &cobra.Command{
// 	Use:   "config",
// 	Short: "Manage pingctl configuration to connect to a PingOne tenant",
// 	Long:  `Provide connection details through which the pingctl can authenticate to a running PingOne tenant.`,
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		if err := pingoneConfig.ConfigurePingOneProfile(cmd.Context()); err != nil {
// 			return err
// 		}

// 		return nil
// 	},
// }

// func init() {
// 	// Here you will define your flags and configuration settings.

// 	// Cobra supports Persistent Flags which will work for this command
// 	// and all subcommands, e.g.:
// 	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	// Cobra supports local flags which will only run when this command
// 	// is called directly, e.g.:
// 	// ConfigCmd.Flags().StringVar(&pingoneConfig.EnvironmentId, "environmentID", "", "Environment ID of the PingOne tenant")
// 	// ConfigCmd.Flags().StringVar(&pingoneConfig.ClientId, "clientID", "", "Client ID of the PingOne tenant to authenticate with")
// }
