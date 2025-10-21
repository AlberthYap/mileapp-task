package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database

func ConnectDB(uri string) {
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  client, err := mongo.Connect(options.Client().ApplyURI(uri))
  if err != nil {
    log.Fatal("Failed to connect to MongoDB:", err)
  }

  // Ping database
  err = client.Ping(ctx, nil)
  if err != nil {
    log.Fatal("Failed to ping MongoDB:", err)
  }

  // Get database name from env
  dbName := os.Getenv("DB_NAME")
  if dbName == "" {
    dbName = "task_management"
  }

  Client = client
  DB = client.Database(dbName)
  
  log.Println("Connected to MongoDB!")
  log.Printf("Database: %s\n", dbName)
}
