package config

import (
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Port int

	QdbConfig      *QdbConfig      `mapstructure:"qdb"`
	InfluxDBConfig *InfluxDBConfig `mapstructure:"influxdb"`
}

type QdbConfig struct {
	Host       string
	Port       int
	InfluxPort int
	User       string
	Password   string
	Dbname     string
	SslMode    string
}

type InfluxDBConfig struct {
	Host         string
	Token        string
	Organization string
	Bucket       string
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
	viper.SetDefault("influxdb.host", "http://localhost:8086")
	viper.SetDefault("influxdb.token", "")
	viper.SetDefault("influxdb.organization", "wireless-monkeys")
	viper.SetDefault("influxdb.bucket", "people_count")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	viper.Unmarshal(&config)
}
