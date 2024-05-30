package configuration

type AppConfig struct {
	Database DatabaseConfig `yaml:"database"`
}

type DatabaseConfig struct {
	Driver         string `yaml:"driver"`
	Host           string `yaml:"host"`
	Name           string `yaml:"name"`
	Password       string `yaml:"password"`
	Port           int    `yaml:"port"`
	User           string `yaml:"user"`
	DefaultTimeout int    `yaml:"default_timeout"`
}
