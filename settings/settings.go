package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

// init settings
func Init() (err error) {
	// set config file info
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	// start to read config
	err = viper.ReadInConfig()

	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed, err: %v\n", err)
		return

	}

	// watch config change
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("The config file has been modified")
	})
	return
}
