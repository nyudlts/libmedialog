package main

import postgres_driver "libmedialog/postgres"

func main() {
	if err := postgres_driver.InitDB("dbconfig.yml", "localhost"); err != nil {
		panic(err)
	}
}
