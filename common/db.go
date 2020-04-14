package common

import (
	"os"

	mgo "gopkg.in/mgo.v2"
)

// DBConnect defines the connection structure
type DBConnect struct {
	session *mgo.Session
}

// ConnectDB is setup db
func ConnectDB() (conn *DBConnect) {
	host := os.Getenv("MONGO_HOST")
	if host == "" {
		host = "localhost:27017"
	}

	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	conn = &DBConnect{session}

	return conn
}

// Use handles connect to a certain collection
func (conn *DBConnect) Use(tbName string) (collection *mgo.Collection) {
	dbName := os.Getenv("MONGO_DB_NAME")
	if dbName == "" {
		dbName = "gin-cbot"
	}

	// This returns method that interacts with a specific collection and table
	return conn.session.DB(dbName).C(tbName)
}

// Close handles closing a database connection
func (conn *DBConnect) Close() {
	// This closes the connection
	conn.session.Close()
	return
}
