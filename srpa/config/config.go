package config

import (
	"fmt"
	"log"
	"os"

	"github.com/or1ko/srpa/srpa/logging"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ReverseUrl string                `yaml:"reverse-url"`
	Port       string                `yaml:"port"`
	Mail       MailConfig            `yaml:"mail"`
	Logging    logging.LoggingConfig `yaml:"logging"`
}

type MailConfig struct {
	Host        string   `yaml:"host"`
	From        string   `yaml:"from"`
	User        string   `yaml:"user"`
	Pass        string   `yaml:"pass"`
	MailAddress []string `yaml:"validedlist"`
}

func Load(filename string) Config {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	if err := yaml.Unmarshal(bytes, &config); err != nil {
		log.Fatal(err)
	}

	return config
}

func (config Config) Save(filename string) {
	yaml, err := yaml.Marshal(config)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(yaml)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
