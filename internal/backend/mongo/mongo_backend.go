package mongo

import (
	"context"
	"github.com/turispro/turislog/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var client *mongo.Client

type MongoBackend struct {
}

func (mb MongoBackend) Register(message model.Log) {
	coll := client.Database("turispro_logs").Collection(os.Getenv("ELASTICSEARCH_INDEX"))
	if _, err := coll.InsertOne(context.TODO(), message); err != nil {
		log.Println(err)
	}
}

func NewMongoBackend() MongoBackend {
	client = start()
	return MongoBackend{}
}

func start() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		panic(err)
	}
	return client
}
