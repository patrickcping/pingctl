package generate

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Tools to generate/export configuration files from the Ping platform",
	Long: `Tools to connect to and manage a PingOne tenant.  This includes the ability to authenticate to a PingOne tenant and generate configuration files for PingOne resources.

	Examples:
	
		pingctl generate pingone --terraform --environmentID <environmentID> --output <output directory>
	
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("generate called")
		return nil
	},
}

func init() {

	GenerateCmd.AddCommand(connector.GenerateConnectorCommands()...)
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
