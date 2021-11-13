package core

import (
	"CommonConfig/global"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		config = "config.toml"
		fmt.Printf("您未手动设置配置文件,使用默认配置[%s]\n", config)
	} else {
		config = path[0]
		fmt.Printf("您已手动设置配置文件[%s]\n", config)
	}
	fmt.Println("...")
	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(
		func(e fsnotify.Event) {
			fmt.Println("config file changed:", e.Name)
			if err := v.Unmarshal(&global.CONFIG); err != nil {
				fmt.Println(err)
			}
		},
	)
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("配置文件[%s]加载成功\n", config)
	return v
}
