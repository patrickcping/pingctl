package connector

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GenerateConnectorCommands() []*cobra.Command {
	connectors := Connectors()

	commands := make([]*cobra.Command, 0)

	for _, connector := range connectors {

		providerCanGenerate := false

		connectorWithTerraformProvider, connectorIsTerraformProvider := connector().(CmdConnectorWithTerraformProvider)
		if connectorIsTerraformProvider {
			providerCanGenerate = true
		}

		_, connectorIsCustomExport := connector().(CmdConnectorWithCustomExport)
		if connectorIsCustomExport {
			providerCanGenerate = true
		}

		if !providerCanGenerate {
			continue
		}

		cobraCommand := &cobra.Command{
			Use:   connector().CommandName(),
			Short: "Tools to generate/export configuration files from " + connector().ProductName() + ".",
			Long: `		
			Examples:
			
				pingctl generate ` + connector().CommandName() + ` --terraform --environmentID <environmentID> --output <output directory>
			
			`,
			RunE: func(cmd *cobra.Command, args []string) error {

				if err := connector().ConfigureConnector(cmd.Context(), "dev"); err != nil {
					return err
				}

				if err := connector().TestConnection(cmd.Context()); err != nil {
					return err
				}

				if err := connectorWithTerraformProvider.GenerateTerraformHCL(cmd.Context()); err != nil {
					return err
				}

				return nil
			},
		}

		if connectorIsTerraformProvider {
			cobraCommand.Flags().BoolP("terraform", "t", true, "Generate Terraform output.")
		}

		if connectorIsCustomExport {
			defaultValue := false
			if !connectorIsTerraformProvider {
				defaultValue = true
			}
			cobraCommand.Flags().BoolP("custom", "c", defaultValue, "Generate custom output.")
		}

		cobraCommand.PersistentFlags().String("adminClientID", "", "The admin client ID.")
		cobraCommand.MarkFlagRequired("adminClientID")
		viper.BindPFlag("profiles.default.pingone.adminClientId", cobraCommand.PersistentFlags().Lookup("adminClientID"))

		commands = append(commands, cobraCommand)
	}

	return commands
}
