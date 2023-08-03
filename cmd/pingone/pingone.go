package pingone

// import (
// 	"github.com/pingidentity/pingctl/cmd/pingone/config"
// 	"github.com/pingidentity/pingctl/cmd/pingone/generate"
// 	"github.com/pingidentity/pingctl/cmd/pingone/login"
// 	"github.com/pingidentity/pingctl/cmd/pingone/profile"
// 	"github.com/spf13/cobra"
// 	"github.com/spf13/viper"
// )

// // pingoneCmd represents the pingone command
// var PingoneCmd = &cobra.Command{
// 	Use:   "pingone",
// 	Short: "Tools to connect to and manage a PingOne tenant",
// 	Long: `Tools to connect to and manage a PingOne tenant.  This includes the ability to authenticate to a PingOne tenant and generate configuration files for PingOne resources.

// Examples:
//   pingctl pingone config
//   pingctl pingone generate --terraform --environmentID <environmentID> --output <output directory>
// `,
// }

// func init() {
// 	PingoneCmd.AddCommand(config.ConfigCmd)
// 	PingoneCmd.AddCommand(generate.GenerateCmd)
// 	PingoneCmd.AddCommand(login.LoginCmd)
// 	PingoneCmd.AddCommand(profile.ProfileCmd)

// 	// Here you will define your flags and configuration settings.

// 	// Cobra supports Persistent Flags which will work for this command
// 	// and all subcommands, e.g.:
// 	// pingoneCmd.PersistentFlags().String("foo", "", "A help for foo")
// 	PingoneCmd.PersistentFlags().String("adminClientID", "", "The admin client ID.")
// 	PingoneCmd.MarkFlagRequired("adminClientID")
// 	viper.BindPFlag("profiles.default.pingone.adminClientId", PingoneCmd.PersistentFlags().Lookup("adminClientID"))

// 	PingoneCmd.PersistentFlags().String("adminClientSecret", "", "The admin client secret.")
// 	PingoneCmd.MarkFlagRequired("adminClientSecret")
// 	viper.BindPFlag("profiles.default.pingone.adminClientSecret", PingoneCmd.PersistentFlags().Lookup("adminClientSecret"))

// 	PingoneCmd.PersistentFlags().String("adminEnvironmentID", "", "The admin client environment ID.")
// 	PingoneCmd.MarkFlagRequired("adminEnvironmentID")
// 	viper.BindPFlag("profiles.default.pingone.adminEnvironmentId", PingoneCmd.PersistentFlags().Lookup("adminEnvironmentID"))

// 	PingoneCmd.PersistentFlags().String("apiHostname", "api.pingone.eu", "The admin client API hostname.")
// 	PingoneCmd.MarkFlagRequired("apiHostname")
// 	viper.BindPFlag("profiles.default.pingone.apiHostname", PingoneCmd.PersistentFlags().Lookup("apiHostname"))

// 	PingoneCmd.PersistentFlags().String("authHostname", "auth.pingone.eu", "The admin client auth hostname.")
// 	PingoneCmd.MarkFlagRequired("authHostname")
// 	viper.BindPFlag("profiles.default.pingone.authHostname", PingoneCmd.PersistentFlags().Lookup("authHostname"))

// 	// Cobra supports local flags which will only run when this command
// 	// is called directly, e.g.:
// 	// pingoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
