package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"mailerservice/pkg/utils"
)

// EnvConfigs is a global variable that contains all environment variables
var EnvConfigs *envConfigs

// InitEnvConfigs initializes the EnvConfigs variable
func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}

// envConfigs is a struct that contains all environment variables
type envConfigs struct {
	LocalServerPort      string `mapstructure:"LOCAL_SERVER_PORT"`
	SmtpHost             string `mapstructure:"SMTP_HOST"`
	SmtpPort             int    `mapstructure:"SMTP_PORT"`
	MailerSenderName     string `mapstructure:"MAILER_SENDER_NAME"`
	MailerSenderAddress  string `mapstructure:"MAILER_SENDER_ADDRESS"`
	MailerSenderPassword string `mapstructure:"MAILER_SENDER_PASSWORD"`
}

// loadEnvVariables loads all environment variables from the dev.env file
func loadEnvVariables() *envConfigs {
	// Tell the viper the path/location of the configuration file
	viper.AddConfigPath(".")
	// Tell viper the name of the configuration file (without the extension)
	viper.SetConfigName("dev")
	// Tell viper the configuration type
	viper.SetConfigType("env")

	// Read the configuration file
	err := viper.ReadInConfig()
	utils.FatalErr("Error reading config file! ", err)

	// Unmarshal the configuration file into a struct
	var config envConfigs
	err = viper.Unmarshal(&config, func(c *mapstructure.DecoderConfig) {
		c.TagName = "mapstructure"
	})
	utils.FatalErr("Unable to decode into struct! ", err)

	return &config
}
