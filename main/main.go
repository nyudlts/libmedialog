package main

import (
	"gorm.io/gorm"
	mongoDriver "libmedialog/mongo"
	postgresDriver "libmedialog/postgres"
)

var pgClient *gorm.DB

func main() {
	if err := postgresDriver.InitDB("dbconfig.yml", "localhost"); err != nil {
		panic(err)
	}

	if err := mongoDriver.InitMongo(); err != nil {
		panic(err)
	}

	pgClient = postgresDriver.MedialogDB

	err := mongoDriver.InsertAllIntoMongoDB(pgClient)
	if err != nil {
		panic(err)
	}

}
