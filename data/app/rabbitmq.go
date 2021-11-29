package app

import (
	"bytes"
	"fmt"
)

type RabbitMQConfig struct {
	CMS      RabbitConfig `yaml:"cms"`
	API      RabbitConfig `yaml:"api"`
	Consumer string       `yaml:"consumer"`
}

type RabbitConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Vhost    string `yaml:"vhost"`
}

func (config RabbitConfig) Addr() string {
	var buf bytes.Buffer

	buf.WriteString("amqp://")
	buf.WriteString(config.Username)
	buf.WriteString(":")
	buf.WriteString(config.Password)

	buf.WriteString(fmt.Sprintf("@%v:%v/%v", config.Host, config.Port, config.Vhost))
	url := buf.String()
	return url
}
