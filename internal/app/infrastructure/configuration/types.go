package configuration

type AppConfig struct {
	Database DatabaseConfig `yaml:"database"`
	Token    TokenConfig    `yaml:"token"`
	Premium  PremiumConfig  `yaml:"premium"`
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

type TokenConfig struct {
	DefaultExpiration int    `yaml:"default_expiration"`
	Key               string `yaml:"jwt_key"`
}

type PremiumConfig struct {
	MapPackageToID map[string]int `yaml:"mapPackagetoID"`
}
