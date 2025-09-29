package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server" env-required:"true"`
}

type HTTPServer struct {
	Address string `yaml:"address" env-default:"localhost"`
	PORT    string `yaml:"port" env-default:"8080"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "path to configuration file")
		flag.Parse()
		configPath = *flags
		if configPath == "" {
			log.Fatal("Config Path is not Set!")
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config File doesnot exists %s", configPath)
	}
	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Cannot read the Config File: %s", err.Error())
	}
	return &cfg
}
