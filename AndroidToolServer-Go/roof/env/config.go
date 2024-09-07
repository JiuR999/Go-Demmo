package env

import (
	"AndroidToolServer-Go/common"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var Config = initConfig()

type config struct {
	DBName     string
	DBUserName string
	DBPassword string
	DBHost     string
	DBPort     string
	ServerPort string
}

func initConfig() *config {
	hit := false
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error loading pwd dir:", err.Error())
	} else {
		if FileExists(pwd + "/" + common.CONFIG_FILE_NAME) {
			err := godotenv.Overload(common.CONFIG_FILE_NAME)
			if err != nil {
				fmt.Println("Error overloading conf.env file:", err.Error())
			} else {
				hit = true
			}
		}
	}

	if hit == false {
		fmt.Println("can not find any conf.env file")
		panic("can not find any conf.env file")
	}

	envConfig := &config{
		DBName:     getOrDefault("DB_Name", "dgth"),
		DBHost:     getOrDefault("DB_HOST", "127.0.0.1"),
		DBUserName: getOrDefault("DB_USERNAME", "root"),
		DBPassword: getOrDefault("DB_PASSWORD", "123456"),
		DBPort:     getOrDefault("DB_PORT", "3306"),
		ServerPort: getOrDefault("SERVER_PORT", "8080"),
	}
	return envConfig
}

func getOrDefault(key string, def string) string {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	return value
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return true
	}
	return true
}
