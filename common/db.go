package common

import (
	"fmt"
	"os"
	"time"

	"github.com/Kamva/mgm/v2"
	"github.com/polnoy/gin-cbot/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// ConnectDb defined connnect database
func ConnectDb() {
	host := os.Getenv("MONGO_HOST")
	if host == "" {
		host = "mongodb://localhost:27017"
	}

	db := os.Getenv("MONGO_DB_NAME")
	if db == "" {
		db = "gin-cbot"
	}

	// Setup mgm default config
	err := mgm.SetDefaultConfig(&mgm.Config{CtxTimeout: 12 * time.Second}, db, options.Client().ApplyURI(host))
	if err != nil {
		fmt.Println("Connect database error: ", err)
	}

	// Create user indexes
	user := mgm.Coll(&model.User{})
	if _, err := user.Indexes().CreateMany(
		mgm.Ctx(),
		[]mongo.IndexModel{
			{
				Keys:    bson.M{"email": 1},
				Options: options.Index().SetUnique(true),
			},
		},
	); err != nil {
		fmt.Println("User index error: ", err)
	}
}
