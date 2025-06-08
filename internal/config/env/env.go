package env

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	errs "github.com/Saidurbu/go-qatarat-dashboard/internal/domain"
)

const defaultEnvFileName = "development"

type Environment string

const (
	EnvironmentDevelopment Environment = "development"
	EnvironmentProduction  Environment = "production"
	EnvironmentStaging     Environment = "staging"
	EnvironmentTest        Environment = "test"
	EnvironmentDocker      Environment = "docker"
)

type Env struct {
	Environment                     Environment `json:"environment" validate:"required,oneof=development production staging test docker"`
	AppPort                         int         `json:"appPort" validate:"required"`
	DatabaseURL                     string      `json:"databaseURL" validate:"required"`
	RedisDatabaseURL                string      `json:"redisDatabaseURL" validate:"required"`
	JWTClientAccessTokenSecretKey   string      `json:"jWTClientAccessTokenSecretKey" validate:"required"`
	JWTMerchantAccessTokenSecretKey string      `json:"jWTMerchantAccessTokenSecretKey" validate:"required"`
	JWTAdminAccessTokenSecretKey    string      `json:"jWTAdminAccessTokenSecretKey" validate:"required"`
	DBDriver                        string      `json:"databaseDriver"`
	MyFatoorahApiKey                string      `json:"myFatoorahApiKey"`
	ApplicationUrl                  string      `json:"applicationUrl"`
}

func NewEnv() *Env {

	e := &Env{}
	if err := e.loadEnv(); err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}
	log.Println("failed to load environment variables: %v", e)
	return e

}

func (e *Env) loadEnv() error {
	bytes, err := e.getEnvFile()
	if err != nil {
		log.Fatalf("%v", err)
	}

	if err := json.Unmarshal(bytes, &e); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	return nil

}

func (e *Env) getEnvFile() (envFile []byte, err error) {
	var envFileName string
	environment := os.Getenv("ENVIRONMENT")

	log.Printf("Application running on: %s environment", environment)

	if environment != "" {
		envFileName = fmt.Sprintf("env.%s.json", strings.ToLower(environment))
	} else {
		envFileName = fmt.Sprintf("env.%s.json", defaultEnvFileName)
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("cannot get caller info")
	}

	envFilePath := filepath.Join(filepath.Dir(filename), envFileName)

	envFile, err = os.ReadFile(envFilePath)
	if err != nil {
		return nil, errs.New(err)
	}

	return envFile, nil
}
