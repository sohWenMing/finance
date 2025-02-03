package loadenv

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(path string) (err error) {
	err = godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}

func GetEnvVar(envPath string, envKey string) (envVar string, err error) {
	err = LoadEnv(envPath)
	if err != nil {
		return "", err
	}
	envVar = os.Getenv(envKey)
	if envVar == "" {
		return "", fmt.Errorf("environment variable with key %s does not exist", envKey)
	}
	return envVar, nil
}
