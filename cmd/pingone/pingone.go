package pingone

import (
	"github.com/pingidentity/pingctl/cmd/pingone/config"
	"github.com/pingidentity/pingctl/cmd/pingone/generate"
	"github.com/pingidentity/pingctl/cmd/pingone/login"
	"github.com/pingidentity/pingctl/cmd/pingone/profile"
	"github.com/spf13/cobra"
)

// pingoneCmd represents the pingone command
var PingoneCmd = &cobra.Command{
	Use:   "pingone",
	Short: "Tools to connect to and manage a PingOne tenant",
	Long: `Tools to connect to and manage a PingOne tenant.  This includes the ability to authenticate to a PingOne tenant and generate configuration files for PingOne resources.

Examples:
  pingctl pingone config
  pingctl pingone generate --terraform --environmentID <environmentID> --output <output directory>
`,
}

func init() {
	PingoneCmd.AddCommand(config.ConfigCmd)
	PingoneCmd.AddCommand(generate.GenerateCmd)
	PingoneCmd.AddCommand(login.LoginCmd)
	PingoneCmd.AddCommand(profile.ProfileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
