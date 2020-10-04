package config

import (
	"log"
	"strings"
)

//Config Structure

type Config struct {
	AppPort   string
	MongoURL  string
	MongoPort string
	MongoAuth string
	LogPath   string
}

//Load function
func Load() Config {
	var configtext Config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	user := viper.GetString("mongo.username")
	pwd := viper.GetString("mongo.password")
	auth := ""
	if len(user) != 0 && len(pwd) != 0 {
		auth = user + ":" + pwd + "@"
	}

	configtext = Config{
		AppPort:   viper.GetString("app.port"),
		MongoURL:  viper.GetString("mongo.url"),
		MongoPort: viper.GetString("mongo.port"),
		MongoAuth: auth,
	}

	return configtext

}
