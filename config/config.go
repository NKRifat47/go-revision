package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration *Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JwtSecretKey string
}

func LoadConfig() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	serviceName := os.Getenv("SERVICE_NAME")
	httpPort := os.Getenv("HTTP_PORT")

	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}
	if serviceName == "" {
		fmt.Println("Service name is required")
		os.Exit(1)
	}
	if httpPort == "" {
		fmt.Println("HTTP port is required")
		os.Exit(1)
	}

	httpPortInt, err := strconv.Atoi(httpPort)
	if err != nil {
		fmt.Println("HTTP port must be an integer")
		os.Exit(1)
	}

	JwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if JwtSecretKey == "" {
		fmt.Println("Jwt Secret Key is required")
		os.Exit(1)
	}

	configuration = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    httpPortInt,
		JwtSecretKey: JwtSecretKey,
	}
}

func GetConfig() *Config {
	if configuration == nil {
		LoadConfig()
	}
	return configuration
}