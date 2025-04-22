package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port       string `json:"port"`
	AWSRegion  string `json:"awsRegion"`
	DBHost     string `json:"dbHost"`
	DBUser     string `json:"dbUser"`
	DBPassword string `json:"dbPassword"`
}

func Load(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := &Config{}
	err = decoder.Decode(cfg)
	return cfg, err
}
