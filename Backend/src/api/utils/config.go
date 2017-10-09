package utils

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

var CONFIG *Config

type Config struct {
	Port    string `yaml:"port"`
}

func LoadConfig(file string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Cannot load config file")
	}
	err = yaml.Unmarshal(data, &CONFIG)
	if err != nil {
		fmt.Println("Cannot parse config file")
	}
}
