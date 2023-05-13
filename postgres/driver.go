package postgres

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type DBConfig struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Port     int    `yaml:"port"`
	SSLMode  string `yaml:"ssl_mode"`
}

var dbConfig *DBConfig
var MedialogDB *gorm.DB

func InitDB(configFile string, environment string) {
	GetDBConfig(configFile, environment)
	var err error
	MedialogDB, err = gorm.Open(postgres.Open(dbConfig.getDSN()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func (dbconfig *DBConfig) getDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", dbconfig.Host, dbconfig.Username, dbconfig.Password, dbconfig.Database, dbconfig.Port, dbconfig.SSLMode)
}

func GetDBConfig(configFile string, environment string) {
	dbConfigMap := map[string]DBConfig{}
	configBytes, err := os.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configBytes, &dbConfigMap)
	if err != nil {
		panic(err)
	}

	dbc := DBConfig{}
	for k, v := range dbConfigMap {
		if environment == k {
			fmt.Println(v.Port)
			dbc.Host = v.Host
			dbc.Port = v.Port
			dbc.SSLMode = v.SSLMode
			dbc.Database = v.Database
			dbc.Username = v.Username
			dbc.Password = v.Password
		}
	}
	dbConfig = &dbc
}
