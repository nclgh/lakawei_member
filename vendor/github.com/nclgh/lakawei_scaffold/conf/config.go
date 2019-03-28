package conf

import (
	"sync"
	"github.com/jinzhu/configor"
)

var config *YamlConfig

func initConfig() {
	data := make(map[interface{}]interface{})
	configor.New(&configor.Config{}).Load(data, "/opt/lakawei/conf/config/config.yml")
	config = &YamlConfig{data}
}

var once sync.Once

func GetConfig() *YamlConfig {
	once.Do(func() { initConfig() })
	return config
}
