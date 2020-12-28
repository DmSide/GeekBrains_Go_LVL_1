package lesson7_cfgReader

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type YamlConfig struct {
	Port        int    `yaml:"port"`
	DbURL       string `yaml:"db_port"`
	JaegerURL   string `yaml:"jaeger_url"`
	SentryURL   string `yaml:"sentry_url"`
	KafkaBroker string `yaml:"kafka_broker"`
	// MyDate	Date TODO: Add date
}

func ConfigInitByYamlParams() YamlConfig {
	file, err := ioutil.ReadFile("config/conf.yaml")
	if err != nil {
		log.Fatal(err)
	}

	cfg := YamlConfig{}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("Config uploaded successfully!")
	return cfg
}
