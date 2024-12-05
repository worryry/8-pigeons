package setting

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

func Start() {
	err := ConfigInit()
	if err != nil {
		log.Fatalf("init.setup config err: %v", err)
		return
	}
}

func ConfigInit() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		//var configFileNotFoundError viper.ConfigFileNotFoundError
		//if errors.As(err, &configFileNotFoundError) {
		//
		//}
		return err
	} else {
		return nil
	}
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetInt32(key string) int32 {
	return viper.GetInt32(key)
}

func GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

func GetUint(key string) uint {
	return viper.GetUint(key)
}

func GetUint32(key string) uint32 {
	return viper.GetUint32(key)
}

func GetUint64(key string) uint64 {
	return viper.GetUint64(key)
}

func GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}
