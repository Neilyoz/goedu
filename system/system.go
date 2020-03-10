package system

import (
	"gopkg.in/yaml.v2"
	"goedu/utils"
	"io/ioutil"
	"path"
	"sync"
)

// Config 配置结构体
type Config struct {
	Debug bool   `yaml:"debug"`
	Key   string `yaml:"key"`
	Mysql struct {
		User     string `yaml:"user"`
		Pass     string `yaml:"pass"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database string `yaml:"database"`
		Charset  string `yaml:"charset"`
	}
}

var configure *Config
var once sync.Once

// GetConfigure 获取配置结构
func GetConfigure() *Config {
	readConfigure()
	return configure
}

// readConfigure 读取配置文件
func readConfigure() {
	if configure != nil {
		return
	}

	// 保证线程安全
	once.Do(func() {
		filename := path.Join(utils.GetCurrentPath(), "config/config.yaml")
		yamlByte, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		var config Config
		err = yaml.Unmarshal(yamlByte, &config)
		if err != nil {
			panic(err)
		}

		configure = &config
	})
}
