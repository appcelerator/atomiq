package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// Config is for all the configuration settings.
type Config struct {
	Verbose bool
	Github  string
	Target  string
	Images  []string
}

// InitConfig reads in a config file and ENV variables if set.
// Configuration variable lookup occurs in a specific order.
func InitConfig(configFile string, config *Config, verbose bool) {
	config.Verbose = verbose

	// Add matching envirionment variables - will be first in precedence.
	viper.AutomaticEnv()

	// Add config file specified using flag - will be next in precedence.
	if configFile != "" {
		viper.SetConfigFile(configFile)
	}

	// Add default config file (without extension) - will be last in precedence.
	// First search home directory; if not found, then attempt to also search working
	// directory (will only succeed if process was started from application directory).
	viper.SetConfigName(".ampswarm")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")

	// If a config file is found, read it in.
	// Extra check for verbose because it might not have been set by
	// a flag, but might be set in the config
	if err := viper.ReadInConfig(); err == nil {
		if verbose || viper.GetBool("Verbose") {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	} else {
		if verbose || viper.GetBool("Verbose") {
			fmt.Println("warning: no valid configuration file (.ampswarm.yaml) found in home or current directory")
		}
	}

	// Save viper into config
	err := viper.Unmarshal(config)
	if err != nil {
		fmt.Println(err)
		panic("unable to process config")
	}
}

// LoadImageList loads all the images in the config.
func LoadImageList() (images []string, err error) {
	viper.SetConfigName("images")
	viper.AddConfigPath("config")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	images = viper.GetStringSlice("images")
	return
}

// Save saves the configuration to ~/.ampswarm.yaml
func (c *Config) Save() (err error) {
	homedir, err := homedir.Dir()
	if err != nil {
		return
	}
	contents, err := yaml.Marshal(c)
	if err != nil {
		return
	}
	fmt.Println(string(contents))
	err = ioutil.WriteFile(path.Join(homedir, ".ampswarm.yaml"), contents, os.ModeAppend)
	if err != nil {
		return
	}
	return
}
