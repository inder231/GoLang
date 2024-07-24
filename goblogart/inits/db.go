package inits

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var mongoClient *mongo.Client

func MongoDBInit() error {

	mongoURI := os.Getenv("MONGODB_URI")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)
	
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)

    // mongoClient = client
	
	return err
}

var DB *gorm.DB

func MySqlDBInit() {
	dsn := os.Getenv("MYSQL_DB_URI")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to MySQL Database!")
	}

    DB = db

}