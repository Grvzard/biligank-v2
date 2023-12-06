package crud

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Client *mongo.Client
var Coll *mongo.Collection
var Sqldb *gorm.DB

func InitDatabase() {
	uri := os.Getenv("API_MONGO_URI")
	if uri == "" {
		panic("empty env: API_MONGO_URI")
	}
	opts := options.Client().ApplyURI(uri)
	Client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	Coll = Client.Database(os.Getenv("API_MONGO_DB")).Collection(os.Getenv("API_MONGO_COLL"))

	sql_dsn := os.Getenv("API_SQL_DSN")
	Sqldb, err = gorm.Open(mysql.Open(sql_dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func DestoryDatabase() {
	if Client != nil {
		if err := Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}
}
