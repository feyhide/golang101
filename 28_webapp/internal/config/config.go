package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	HTTPServer  HTTPServer `yaml:"http_server"`
}

// this must execute in order to spin up the server
// func with must does not return err as it is a must log error within
// this func and exit

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	// if config path is there ???
	if configPath == "" {

		//if not in env , then check whether passed in arg
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Warning: Config path is not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Warning: Config file does not exist: %s", configPath)
	}

	var cfg Config

	// ReadConfig loads configuration from the specified file into the cfg struct.
	// If the file can't be read or has invalid data, it logs a fatal error and exits.
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Warning: Cannot read config file: %s", err.Error())
	}

	return &cfg
}
