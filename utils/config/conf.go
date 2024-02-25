package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var GlobalConfig Config // 定義全局變量來存儲配置
type Config struct {
	Database struct {
		SQLUrl string `yaml:"sql_url"`
	} `yaml:"database"`
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
}

func (c *Config) GetDatabaseUrl() string {
	return c.Database.SQLUrl
}

func (c *Config) GetServerPort() int {
	return c.Server.Port
}

func init() {
	// 在包初始化時讀取並解析配置文件
	err := ReadConfig("conf/config.yml")
	if err != nil {
		panic(fmt.Errorf("無法讀取配置文件：%v", err))
	}
}

func ReadConfig(filename string) error {
	// 讀取配置文件
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("無法讀取配置文件：%v", err)
	}

	// 解析配置文件並存儲到全局變量中
	err = yaml.Unmarshal(data, &GlobalConfig)
	if err != nil {
		return fmt.Errorf("無法解析配置文件：%v", err)
	}

	return nil
}
