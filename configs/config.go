/*
* parse the project all configs
* A clean way to config your project
* referrenc https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64
* author :tf
 */
package configs

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
	} `yaml:"database"`
}
