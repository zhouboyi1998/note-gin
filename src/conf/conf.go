package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Configuration struct {
	Server struct {
		Host string
		Port string
	}
	Mongo struct {
		Host       string
		Port       string
		Username   string
		Password   string
		Database   string
		Collection string
	}
}

var Config = &Configuration{}

func init() {
	// os.Getenv() 获取环境变量, ioutil.ReadFile() 读取对应的 yaml 配置文件
	file, err := ioutil.ReadFile("src/conf/" + os.Getenv("ENVCONFIG") + ".yaml")
	if err != nil {
		log.Println(err)
	}

	// yaml.Unmarshal() 将 .yaml 配置文件中的配置解析到 Configuration 类型的变量中
	err = yaml.Unmarshal(file, Config)
	if err != nil {
		log.Println(err)
	}
}
