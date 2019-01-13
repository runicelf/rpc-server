package config

import (
	"fmt"
	"github.com/runicelf/rpc-server/models"
	"os"
)

func Get() models.Config {
	return models.Config{
		DBUser:     getEnvVarValue("DB_USER"),
		DBName:     getEnvVarValue("DB_NAME"),
		DBPassword: getEnvVarValue("DB_PASSWORD"),
		DriverName: getEnvVarValue("DB_DRIVER_NAME"),
	}
}

func getEnvVarValue(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Errorf("envVar %s not found", key))
	}
	return value
}
