package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration *Config

type DBConfig struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	EnableSSLMode bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB           *DBConfig
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

	host := os.Getenv("DB_HOST")
	if host == "" {
		fmt.Println("Host is required")
		os.Exit(1)
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		fmt.Println("Port is required")
		os.Exit(1)
	}

	dbPrt, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("DB port must be an integer")
		os.Exit(1)
	}

	name := os.Getenv("DB_NAME")
	if name == "" {
		fmt.Println("Name is required")
		os.Exit(1)
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		fmt.Println("DB User is required")
		os.Exit(1)
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		fmt.Println("DB Password is required")
		os.Exit(1)
	}

	enableSSLModeStr := os.Getenv("DB_ENABLE_SSL_MODE")
	if enableSSLModeStr == "" {
		fmt.Println("ENABLE_SSL_MODE is required")
		os.Exit(1)
	}

	enableSSLModeBool, err := strconv.ParseBool(enableSSLModeStr)
	if err != nil {
		fmt.Println("ENABLE_SSL_MODE must be a boolean (true/false)")
		os.Exit(1)
	}

	configuration = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     httpPortInt,
		JwtSecretKey: JwtSecretKey,
		DB: &DBConfig{
			Host:          host,
			Port:          dbPrt,
			Name:          name,
			User:          user,
			Password:      password,
			EnableSSLMode: enableSSLModeBool,
		},
	}
}

func GetConfig() *Config {
	if configuration == nil {
		LoadConfig()
	}
	return configuration
}
