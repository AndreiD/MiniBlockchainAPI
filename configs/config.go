package configs

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

// Load the config files
func Load() *viper.Viper {
	// Configs
	Config, err := readConfig("/api_config", map[string]interface{}{
		"port":         9090,
		"hostname":     "localhost",
		"environment":  "debug",
		"node_address": "http://localhost:8545",
	})
	if err != nil {
		log.Errorf("error when reading config: %v\n", err)
	}
	return Config
}

//read the config file, helper function
func readConfig(filename string, defaults map[string]interface{}) (*viper.Viper, error) {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}
	v.SetConfigName(filename)
	v.AddConfigPath("./configs")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	return v, err
}
