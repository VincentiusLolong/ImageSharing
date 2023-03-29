package conf

import (
	"fmt"
	"mestorage/models"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("app")
	viper.AddConfigPath("./conf")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	models.RedisSetting.Host = viper.GetString("redis.Host")
	models.RedisSetting.Password = viper.GetString("redis.Password")
	models.RedisSetting.MaxIdle = viper.GetInt("redis.MaxIdle")
	models.RedisSetting.MaxActive = viper.GetInt("redis.MaxActive")
	models.RedisSetting.IdleTimeout = viper.GetDuration("redis.IdleTimeout")
	models.RedisSetting.DB = viper.GetInt("redis.DB")
}
