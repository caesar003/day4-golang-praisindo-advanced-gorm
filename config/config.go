package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Credentials struct {
	DBHost string `json:"dbHost"`
	DBPort string `json:"dbPort"`
	DBUser string `json:"dbUser"`
	DBPass string `json:"dbPass"`
	DBName string `json:"dbName"`
}

func LoadCredentials(filePath string) (*Credentials, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %v", err)
	}

	var creds Credentials
	err = json.Unmarshal(file, &creds)
	if err != nil {
		return nil, fmt.Errorf("unable to parse credentials: %v", err)
	}

	return &creds, nil
}
