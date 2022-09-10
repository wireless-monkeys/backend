package config

import (
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Port int

	QdbConfig *QdbConfig `mapstructure:"qdb"`
}

type QdbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
	SslMode  string
}

var config = Config{}
var once sync.Once

func NewConfig() *Config {
	once.Do(parseConfig)
	return &config
}

func parseConfig() {
	viper.SetDefault("port", 4000)
	viper.SetDefault("qdb.host", "localhost")
	viper.SetDefault("qdb.port", 8812)
	viper.SetDefault("qdb.user", "admin")
	viper.SetDefault("qdb.password", "quest")
	viper.SetDefault("qdb.dbname", "qdb")
	viper.SetDefault("qdb.sslMode", "disable")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.ReadInConfig()
	viper.Unmarshal(&config)
}
