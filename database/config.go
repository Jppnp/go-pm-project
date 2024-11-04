package database

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

var GlobalConfig *Config

type Config struct {
	Database struct {
		Driver   string `json:"driver"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
		SSLMode  string `json:"sslmode"`
		Timezone string `json:"timezone"`
	} `json:"database"`
}

func (c *Config) DSN() string {
	var dsn strings.Builder
	dsn.WriteString(fmt.Sprintf("user=%s dbname=%s sslmode=%s", c.Database.User, c.Database.DBName, c.Database.SSLMode))
	return dsn.String()
}

func LoadConfig() error {
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		return err
	}
	GlobalConfig = &config
	return nil
}
