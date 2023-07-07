/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package generate

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// generateCmd represents the generate command
var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate configuration from a PingOne tenant",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	GenerateCmd.Flags().Bool("terraform", true, "Output in Terraform HCL format")
	GenerateCmd.Flags().StringP("output", "o", "", "The output directory to write the configuration to")
	GenerateCmd.MarkFlagRequired("output")
	GenerateCmd.Flags().StringP("environmentID", "e", "", "The environment ID to generate configuration files for.  If left blank, the default environment ID from the active profile is used.  If there is no default environment ID set in the profile, the command will prompt for the environment ID.")
	viper.BindPFlag("environmentID", GenerateCmd.PersistentFlags().Lookup("environmentID"))
}
