package lesson7_cfgReader

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
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
	file, err := os.Open("./config/conf.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cfg := YamlConfig{}

	err = yaml.NewDecoder(file).Decode(cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("Config uploaded successfully!")
	return cfg
}
