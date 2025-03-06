package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

// Get path of config file
func getConfigPath() (string, error) {
	// Get project directory
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	cfgPath := dir + "/" + configFileName
	return cfgPath, nil
}

// Read .gatorignore file and return it as a Config struct
func Read() (Config, error) {

	// Get config file path
	path, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}

	// Read config file
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	// Unmarshal data into struct
	cfg := Config{}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}
	
	return cfg, nil
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	err := c.save()
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) save() error {
	// Marshal the struct to JSON
	data, err := json.MarshalIndent(c, "", "	")
	if err != nil {
		return err
	}

	// Get config file path
	path, err := getConfigPath()
	if err != nil {
		return err
	}

	// Write data to file
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

