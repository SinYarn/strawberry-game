package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var Conf *Config // 类型为指向 Config 结构体的指针, 值为nil

type Config struct {
	httpPort int
}

// InitConfig  加载配置
func InitConfig(configFile string) {
	Conf = new(Config)
	v := viper.New()
	v.SetConfigFile(configFile)
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		log.Println("Config file changed:", in.Name)
	})
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Read config error, err%v \n", err))
	}

	//解析
	err = v.Unmarshal(&Conf)
	if err != nil {
		panic(fmt.Errorf("Unmarshal config data, err%v \n", err))
	}

}
