package config

import (
	"log"
	"os"
)

type Config struct {
	PublicUrl string
	Ip        string
	Port      string
	DestIp    string
	DestPort  string
}

var GlobalConfig = new(Config)

func (config *Config) SetUpConfig() error {
	config.PublicUrl = os.Getenv("PUBLIC_URL_COMM")
	if config.PublicUrl == "" {
		log.Fatal("missing enviroment variable PUBLIC_URL_COMM")
	}
	config.Ip = os.Getenv("SERVER_IP_COMM")
	if config.Ip == "" {
		config.Ip = "127.0.0.1"
	}
	config.Port = os.Getenv("SERVER_PORT_COMM")
	if config.Port == "" {
		config.Port = "443"
	}
	config.DestIp = os.Getenv("DEST_IP_COMM")
	if config.DestIp == "" {
		log.Fatal("missing enviroment variable DEST_IP_COMM")
	}
	config.DestPort = os.Getenv("DEST_PORT_COMM")
	if config.DestPort == "" {
		log.Fatal("missing enviroment variable DEST_PORT_COMM")
	}
	return nil
}
