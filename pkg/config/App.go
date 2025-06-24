package config

// Book-management-system\pkg\config\App.go

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	db     *mongo.Database
)

func Connect() *mongo.Database {
	if db != nil {
		return db
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database("bookstore")
	log.Println("Connected to MongoDB!")
	return db
}

func GetDB() *mongo.Database {
	if db == nil {
		return Connect()
	}
	return db
}

func CloseDB() {
	if client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}
}
