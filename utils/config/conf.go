package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

var GlobalConfig Config // 定義全局變量來存儲配置
type Config struct {
	Database struct {
		SQLUrl string `yaml:"sql_url"`
	} `yaml:"database"`
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

func InitRootFolder(path string) error {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Join(filepath.Dir(filename), path)
	err := os.Chdir(dir)
	if err != nil {
		return err
	}
	return nil
}
func (c *Config) GetDatabaseUrl() string {
	if envHost := os.Getenv("MYSQL_HOST"); envHost != "" {
		c.Database.SQLUrl = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"),
			os.Getenv("MYSQL_PORT"),
			os.Getenv("MYSQL_DATABASE"))
	}
	return c.Database.SQLUrl
}

func (c *Config) GetServerPort() string {
	return c.Server.Port
}

func InitConfig() {
	absPath, _ := filepath.Abs("conf/config.yml")
	err := ReadConfig(absPath)
	if err != nil {
		panic(fmt.Errorf("config read fail：%v", err))
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
