package configs

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Port string `json:"port"`
		Mode string `json:"mode"`
	} `json:"server"`
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
	} `json:"database"`
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := os.Getenv("APP_ENVIROMENT")
	fmt.Println("env in configs file = ", env)
	var configFile string

	switch env {
	case "development":
		configFile = "configs/config.dev.json"
	case "test":
		configFile = "configs/config.test.json"
	case "staging":
		configFile = "configs/config.stage.json"
	default:
		log.Fatal("No environment specified or invalid APP_ENV")
	}

	file, err := os.Open(configFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}

}
