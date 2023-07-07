package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pingidentity/pingctl/cmd/version"
	"github.com/pingidentity/pingctl/internal/pingone"
	"github.com/pingidentity/pingctl/internal/pingone/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Runtime configuration
type Config struct {
	ConfigPath          string
	ConfigFile          string
	CredentialsPath     string
	CredentialsFile     string
	PingOneClientConfig *client.Config
}

// Config profiles configuration - persists to the file in `Config.ConfigPath/Config.ConfigFile`
type ProfilesConfig struct {
	// The version of config for backward compatibility.  If empty, assume the latest version
	ConfigVersion *string `json:"_configVersion,omitempty"`
	// The profiles associated with the user's cli installation.  If empty from file, run a config init
	Profiles []ProfilesConfig `json:"profiles,omitempty"`
}

// An individual config profile - persists to the file in `ConfigPath/ConfigFile` under `ProfilesConfig`
type ProfileConfig struct {
	// The ID of the profile
	Id             string `json:"id"`
	PingOneCmdAuth *pingone.ProfilePingOneCmdAuthConfig
}

var (
	configType       = "json"
	legacyConfigType = "env"
	configHome       string
)

func Init() error {

	osHome, err := os.UserHomeDir()
	cobra.CheckErr(err)

	configHome = filepath.Join(osHome, ".pingidentity")
	configName := "pingctl-config"
	configPath := filepath.Join(configHome, configName)

	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configHome)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {

			migrated, err := tryMigrateLegacyConfig(configPath)
			if err != nil {
				return err
			}

			if !*migrated {
				if err := writeDefaultConfig(configPath); err != nil {
					return err
				}
			}

		} else {
			return fmt.Errorf("Error reading config file: %s", err)
		}
	}

	if viper.GetString("_configVersion") == "dev" {
		fmt.Println("WARNING: You are using a development version of pingctl.  This may be unstable.")
	}

	// Config upgrading goes here

	return nil
}

func tryMigrateLegacyConfig(newConfigPath string) (*bool, error) {

	legacyConfig := viper.New()

	// Try the legacy config file
	legacyConfigName := "config"
	legacyConfigType := "env"
	legacyConfig.SetConfigName(legacyConfigName)
	legacyConfig.SetConfigType(legacyConfigType)
	legacyConfig.AddConfigPath(configHome)
	legacyConfig.AddConfigPath(".")

	if err := legacyConfig.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Can't migrate if there's no legacy config, return false so we can create some defaults
			returnVar := false
			return &returnVar, nil
		} else {
			return nil, fmt.Errorf("Error reading legacy config file: %s", err)
		}
	} else {
		fmt.Printf("Migrating from legacy config file: %s\n", legacyConfigName)

		configMappings := map[string]string{
			"PINGONE_API_URL":                  "profiles.default.pingone.apihostname",
			"PINGONE_AUTH_URL":                 "profiles.default.pingone.authhostname",
			"PINGONE_ENVIRONMENT_ID":           "profiles.default.pingone.adminenvironmentid",
			"PINGONE_WORKER_APP_CLIENT_ID":     "profiles.default.pingone.adminclientid",
			"PINGONE_WORKER_APP_GRANT_TYPE":    "profiles.default.pingone.adminclientgranttype",
			"PINGONE_WORKER_APP_REDIRECT_URI":  "profiles.default.pingone.adminclientredirecturi",
			"PINGONE_WORKER_APP_CLIENT_SECRET": "profiles.default.pingone.adminclientsecret",

			"PING_IDENTITY_ACCEPT_EULA":     "profiles.default.devops.configurationdefaults.accepteula",
			"PING_IDENTITY_DEVOPS_USER":     "profiles.default.devops.credentials.devopsuser",
			"PING_IDENTITY_DEVOPS_KEY":      "profiles.default.devops.credentials.devopskey",
			"PING_IDENTITY_DEVOPS_HOME":     "profiles.default.devops.projectshomepath",
			"PING_IDENTITY_DEVOPS_REGISTRY": "profiles.default.devops.configurationdefaults.dockerregistry",
			"PING_IDENTITY_DEVOPS_TAG":      "profiles.default.devops.configurationdefaults.dockertag",
		}

		// Remap the new config
		for legacyKey, newKey := range configMappings {
			if legacyConfig.IsSet(legacyKey) {
				migratedValue := legacyConfig.GetString(legacyKey)

				if newKey == "profiles.default.pingone.apihostname" {
					migratedValue = strings.ReplaceAll(migratedValue, "https://", "")
					migratedValue = strings.ReplaceAll(migratedValue, "/v1", "")
				}

				if newKey == "profiles.default.pingone.authhostname" {
					migratedValue = strings.ReplaceAll(migratedValue, "https://", "")
				}
				viper.Set(newKey, migratedValue)
			}
		}

		// Carry the additional config over
		for _, legacyProfileKey := range legacyConfig.AllKeys() {
			if configMappings[strings.ToUpper(legacyProfileKey)] == "" {
				viper.Set(fmt.Sprintf("profiles.default.extra.%s", legacyProfileKey), legacyConfig.Get(legacyProfileKey))
			}
		}

		if err := writeDefaultConfig(newConfigPath); err != nil {
			return nil, err
		}

		migrated := true
		return &migrated, nil
	}
}

func writeDefaultConfig(configPath string) error {

	viper.Set("_configVersion", version.GetVersion())
	viper.Set("profiles.default.name", "Default Profile")

	_, err := os.Stat(configPath)
	if !os.IsExist(err) {
		if _, err := os.Create(configPath); err != nil {
			return fmt.Errorf("Error creating config directory: %s", err)
		}
	}
	if err := viper.WriteConfigAs(configPath); err != nil {
		return fmt.Errorf("Error writing config file: %s", err)
	}

	return nil
}
