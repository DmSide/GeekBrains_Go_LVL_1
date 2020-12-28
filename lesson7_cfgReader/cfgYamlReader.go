package lesson7_cfgReader

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type YamlConfig struct {
	Port        int    `yaml:"port" validate:"required"`
	DbURL       string `yaml:"db_url" validate:"required,url"`
	JaegerURL   string `yaml:"jaeger_url" validate:"required,url"`
	SentryURL   string `yaml:"sentry_url" validate:"required,url"`
	KafkaBroker string `yaml:"kafka_broker" validate:"required,url"`
	CorrectURL  string `yaml:"correct_url" validate:"required,url"`
	// MyDate	Date TODO: Add date
}

func ConfigInitByYamlParams() YamlConfig {
	file, err := os.Open("./config/conf.yaml")
	// You can also use:
	// file, err := ioutil.ReadFile("config/conf.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cfg := YamlConfig{}

	err = yaml.NewDecoder(file).Decode(&cfg)
	// You can also use:
	// err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	validate = validator.New()
	err = validate.Struct(cfg)
	if err != nil {
		log.Fatalf("Validation error: %+v", err)
	}

	fmt.Printf("Config uploaded successfully!")
	return cfg
}
