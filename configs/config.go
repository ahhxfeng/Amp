/*
* parse the project all configs
* A clean way to config your project
* referrenc https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64
* author :tf
* file config -> env config -> command line config
 */
package configs

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
	"gopkg.in/yaml.v3"
)

type Config struct {
	// log config
	//LogLevel string `yaml:"log_level"`
	//Logfile string `yaml:"log_file"`

	// logger config
	LogConfig struct {
		LogLevel string `yaml:"log_level"`
		LogFile  string `yaml:"log_file"`
	} `yaml:"log_config"`

	// server config
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`

	// database config
	Database struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		User string `env:"DB_USER"`
		Pass string `env:"DB_PASS"`
		Name string `env:"DB_NAME"`
	} `yaml:"database"`
	// Grafana config
	Grafana struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		User string `env:"GRAFANA_USER"`
		Pass string `env:"GRAFANA_PASS"`
	} `yaml:"grafana"`
}

// gennerate from copilt not use for this time
func NewConfig() *Config {
	return &Config{
		LogConfig: struct {
			LogLevel string `yaml:"log_level"`
			LogFile  string `yaml:"log_file"`
		}{
			LogLevel: "info",
			LogFile:  "logs/app.log",
		},
		Server: struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		}{
			Host: "localhost",
			Port: 8080,
		},
		Database: struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
			User string `env:"DB_USER"`
			Pass string `env:"DB_PASS"`
			Name string `env:"DB_NAME"`
		}{
			Host: "localhost",
			Port: 3452,
		},
		Grafana: struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
			User string `env:"GRAFANA_USER"`
			Pass string `env:"GRAFANA_PASS"`
		}{
			Host: "localhost",
			Port: 4632,
		},
	}
}

// process err
func processError(err error) {
	fmt.Printf("Error: %v", err)
	os.Exit(-1)
}

// load the global config
func LoadConfig(file string) *Config {
	// file -> env -> command line
	cfg := Config{}

	readFile(&cfg, file)
	readEnv(&cfg)

	return &cfg
}

// read config from the config file

func readFile(cfg *Config, file string) {
	f, err := os.Open(file)
	if err != nil {
		processError(err)
	}
	// file close at the func end
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

// read config from env
func readEnv(cfg *Config) {
	err := env.Parse(cfg)
	if err != nil {
		processError(err)
	}

}
