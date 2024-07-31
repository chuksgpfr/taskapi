package pkg

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Configuration struct {
	PostgresDSN       string `mapstructure:"POSTGRES_DSN"`
	GIN_MODE          string `mapstructure:"GIN_MODE"`
	ServerAddress     string `mapstructure:"SERVER_ADDRESS"`
	LoginSymmetricKey string `mapstructure:"LOGIN_SYMMETRIC_KEY"`
}

func environ() map[string]string {
	m := make(map[string]string)
	for _, s := range os.Environ() {
		a := strings.Split(s, "=")
		m[a[0]] = a[1]
	}
	return m
}

func LoadConfig(path string) (config Configuration, err error) {
	if os.Getenv("GO_ENV") == "production" {
		m := environ()
		mByte, _ := json.Marshal(m)
		json.Unmarshal(mByte, &config)
	} else {
		viper.AddConfigPath(path)
		viper.SetConfigFile(".env")

		viper.WatchConfig()

		viper.AutomaticEnv()
		if err = viper.ReadInConfig(); err != nil {
			return
		}

		err = viper.Unmarshal(&config)
	}

	return
}
