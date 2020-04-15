package common

import (
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
)

// DBConnect defines the connection structure
type DBConnect struct {
	session *mgo.Session
}

// DbName is database name
var DbName = "gin-cbot"

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	db := viper.GetString("MONGO_DB_NAME")
	if db != "" {
		DbName = db
	}
}

// ConnectDB is setup db
func ConnectDB() (conn *DBConnect) {
	host := viper.GetString("MONGO_HOST")
	if host == "" {
		host = "localhost:27017"
	}

	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	conn = &DBConnect{session}

	// Create user index
	idx := mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}
	if err := session.DB(DbName).C("users").EnsureIndex(idx); err != nil {
		panic(err)
	}

	return conn
}

// Use handles connect to a certain collection
func (conn *DBConnect) Use(collName string) (collection *mgo.Collection) {
	// This returns method that interacts with a specific collection and table
	return conn.session.DB(DbName).C(collName)
}

// Close handles closing a database connection
func (conn *DBConnect) Close() {
	// This closes the connection
	conn.session.Close()
	return
}
