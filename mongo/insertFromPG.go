package mongo_driver

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"libmedialog"
)

func InsertAllIntoMongoDB(pgClient *gorm.DB) error {
	err := InsertEntriesIntoMongoDB(pgClient)
	if err != nil {
		return (err)
	}

	err = InsertAccessionsIntoMongoDB(pgClient)
	if err != nil {
		return (err)
	}

	err = InsertCollectionsIntoMongoDB(pgClient)
	if err != nil {
		return (err)
	}

	err = InsertUsersIntoMongoDB(pgClient)
	if err != nil {
		return (err)
	}
	return nil
}

func InsertEntriesIntoMongoDB(pgClient *gorm.DB) error {
	mongoCollection := MongoClient.Database("medialog").Collection("entries")
	var data []libmedialog.MlogEntry
	pgClient.Find(&data)
	for _, obj := range data {
		result, err := mongoCollection.InsertOne(context.TODO(), obj)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Inserted docid - %s into collection: %s\n", result.InsertedID, "entrtries")
	}
	return nil
}

func InsertAccessionsIntoMongoDB(pgClient *gorm.DB) error {
	mongoCollection := MongoClient.Database("medialog").Collection("accessions")
	var data []libmedialog.Accession
	pgClient.Find(&data)
	for _, obj := range data {
		result, err := mongoCollection.InsertOne(context.TODO(), obj)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Inserted docid - %s into collection: %s\n", result.InsertedID, "accessions")
	}
	return nil
}

func InsertCollectionsIntoMongoDB(pgClient *gorm.DB) error {
	mongoCollection := MongoClient.Database("medialog").Collection("collections")
	var data []libmedialog.Collection
	pgClient.Find(&data)
	for _, obj := range data {
		result, err := mongoCollection.InsertOne(context.TODO(), obj)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Inserted docid - %s into collection: %s\n", result.InsertedID, "collections")
	}
	return nil
}

func InsertUsersIntoMongoDB(pgClient *gorm.DB) error {
	mongoCollection := MongoClient.Database("medialog").Collection("user")
	var data []libmedialog.User
	pgClient.Find(&data)
	for _, obj := range data {
		result, err := mongoCollection.InsertOne(context.TODO(), obj)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Inserted docid - %s into collection: %s\n", result.InsertedID, "users")
	}
	return nil
}
