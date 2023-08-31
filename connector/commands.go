package connector

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ConnectorType func() CmdConnector

func GenerateConnectorCommands(ctx context.Context, connectors []ConnectorType) []*cobra.Command {
	l := logger.Get()

	commands := make([]*cobra.Command, 0)

	for _, connector := range connectors {
		l.Debug().Str("connector", connector().Metadata().ProductName).Msg("Checking connector for export capability..")

		if exportableConnector, ok := connector().(CmdConnectorWithExport); ok {
			l.Debug().Str("connector", connector().Metadata().ProductName).Msg("Adding exportable connector")

			cobraCommand := &cobra.Command{
				DisableFlagParsing: true,
				Use:                connector().Metadata().CommandName,
				Short:              "Tools to export configuration files from " + connector().Metadata().ProductName + ".",
				Long: `		
			Examples:
			
				pingctl generate ` + connector().Metadata().CommandName + ` --terraform --environmentID <environmentID> --output <output directory>
			
			`,
				RunE: func(cmd *cobra.Command, args []string) error {
					// Initial command processing
					var connectionParams map[string]interface{}
					if err := processCommand(
						cmd,
						args,
						connector().Metadata().ProfileConfigIndex,
						connector().ConnectorSettings(ctx),
						&connectionParams,
					); err != nil {
						return err
					}

					// Configure the connector
					if err := connector().ConfigureConnector(cmd.Context(), "dev", connectionParams); err != nil {
						return err
					}

					// Test the connection
					if err := connector().TestConnection(cmd.Context()); err != nil {
						return err
					}

					// Run the export
					exportTF, err := cmd.Flags().GetBool("terraform")
					if err != nil {
						return err
					}
					l.Debug().Bool("exportTF", exportTF).Msg("Export terraform flag")

					// Get the output directory
					outputDir, err := cmd.Flags().GetString("output-dir")
					if err != nil {
						return err
					}
					l.Debug().Str("outputDir", outputDir).Msg("Output directory flag")

					if err := exportableConnector.Export(cmd.Context(), GenerateHCLOpts{
						OutputDirectoryPath: outputDir,
					}); err != nil {
						return err
					}

					return nil
				},
			}

			cobraCommand.Flags().BoolP("terraform", "t", true, "Generate Terraform output.")

			// cobraCommand.PersistentFlags().String("adminClientID", "", "The admin client ID.")
			// cobraCommand.MarkFlagRequired("adminClientID")
			// viper.BindPFlag("profiles.default.pingone.adminClientId", cobraCommand.PersistentFlags().Lookup("adminClientID"))

			commands = append(commands, cobraCommand)
		} else {
			l.Debug().Str("connector", connector().Metadata().CommandName).Msg("Connector doesn't support configuration export.")
		}
	}

	return commands
}

func processCommand(cmd *cobra.Command, args []string, profileSettingsIndex string, configParms map[string]ConnectorParam, config *map[string]interface{}) error {
	l := logger.Get()

	p.profileSettings = profileSettings{}

	for i := 0; i < reflect.TypeOf(p.profileSettings).NumField(); i++ {
		field := reflect.TypeOf(p.profileSettings).Field(i)
		fieldProfileIndex, ok := field.Tag.Lookup("profile")
		if !ok {
			continue
		}

		l.Debug().Str("name", field.Name).Str("fieldProfileIndex", fieldProfileIndex).Msg("Test Profile settings")

		fieldValue := reflect.ValueOf(&p.profileSettings).Elem().FieldByName(field.Name)
		if fieldValue.CanSet() {

			configIndex := fmt.Sprintf("profiles.default.%s.%s", profileSettingsIndex, fieldProfileIndex)
			profileValue := viper.GetString(configIndex)

			if fieldValue.Kind() == reflect.Ptr {
				fieldValue.Set(reflect.ValueOf(&profileValue))
			}

			if fieldValue.Kind() == reflect.String {
				fieldValue.SetString(profileValue)
			}

			if fieldValue.Kind() == reflect.Bool {
				if v, err := strconv.ParseBool(profileValue); err == nil {
					fieldValue.SetBool(v)
				} else {
					return fmt.Errorf("Error parsing boolean value for %s: %s", configIndex, err)
				}
			}

			if fieldValue.Kind() == reflect.Int || (fieldValue.Kind() == reflect.Ptr && fieldValue.Elem().Kind() == reflect.Int) {
				if v, err := strconv.ParseInt(profileValue, 10, 32); err == nil {
					fieldValue.SetInt(v)
				} else {
					return fmt.Errorf("Error parsing int value for %s: %s", configIndex, err)
				}
			}

			if fieldValue.Kind() == reflect.Float64 || (fieldValue.Kind() == reflect.Ptr && fieldValue.Elem().Kind() == reflect.Float64) {
				if v, err := strconv.ParseFloat(profileValue, 10); err == nil {
					fieldValue.SetFloat(v)
				} else {
					return fmt.Errorf("Error parsing float value for %s: %s", configIndex, err)
				}
			}
		}
	}

	return nil

}
