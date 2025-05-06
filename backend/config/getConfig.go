package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Database struct {
		PSQL struct {
			Host     string `yaml:"Host"`
			Port     int    `yaml:"Port"`
			User     string `yaml:"User"`
			Password string `yaml:"Password"`
			DBName   string `yaml:"DBName"`
			TimeZone string `yaml:"Timezone"`
		} `yaml:"PSQL"`
		Valkey struct {
			Host     string `yaml:"Host"`
			Port     int    `yaml:"Port"`
			DB       int    `yaml:"PSQL"`
			User     string `yaml:"User"`
			Password string `yaml:"Password"`
		} `yaml:"Valkey"`
	} `yaml:"Database"`
}

func GetConfig() *Config {
	var config *Config

	if os.Args[1] == "dev" {
		file, err := os.Open("./config.yaml")
		if err != nil {
			log.Fatalf("Error opening config file: %v", err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatalf("Error closing config file: %v", err)
			}
		}(file)

		// Read the file and parse the YAML content
		yamlParser := yaml.NewDecoder(file)
		err = yamlParser.Decode(&config)
		if err != nil {
			log.Fatalf("Error parsing YAML file: %v", err)
		}
	} else if os.Args[1] == "prod" {
		config.Database.PSQL.Host = os.Getenv("PSQL_HOST")
		config.Database.PSQL.Port, _ = strconv.Atoi(os.Getenv("PSQL_PORT"))
		config.Database.PSQL.User = os.Getenv("PSQL_USER")
		config.Database.PSQL.Password = os.Getenv("PSQL_PASSWORD")
		config.Database.PSQL.DBName = os.Getenv("PSQL_DBNAME")
		config.Database.PSQL.TimeZone = os.Getenv("PSQL_TIMEZONE")
		config.Database.Valkey.Host = os.Getenv("REDIS_HOST")
		config.Database.Valkey.Port, _ = strconv.Atoi(os.Getenv("REDIS_PORT"))
		config.Database.Valkey.DB, _ = strconv.Atoi(os.Getenv("REDIS_DB"))
		config.Database.Valkey.User = os.Getenv("REDIS_USER")
		config.Database.Valkey.Password = os.Getenv("REDIS_PASSWORD")
	} else {
		panic("Invalid environment")
		return nil
	}

	return config
}
