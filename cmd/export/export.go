package export

import (
	"fmt"

	"github.com/pingidentity/pingctl/connector"
	"github.com/pingidentity/pingctl/internal"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

var (
	outputDir string
)

// ExportCmd represents the export command
var ExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Tools to export configuration as code from the Ping platform",
	Long: `Tools to connect to and manage a PingOne tenant.  This includes the ability to authenticate to a PingOne tenant and generate configuration files for PingOne resources.

	Examples:
	
		pingctl export pingone --environment-id <environmentID> --output terraform

		pingctl export pingone --environment-id <environmentID> --output terraform --output-dir .
	
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Export command called")
		return nil
	},
}

func init() {
	l := logger.Get()

	l.Debug().Msg("Adding commands")

	ExportCmd.AddCommand(connector.GenerateConnectorCommands(internal.Connectors())...)

	l.Debug().Msg("Commands added")

	ExportCmd.PersistentFlags().StringVarP(&outputDir, "output-dir", "d", ".", "The directory in which to save the generated files")
}
