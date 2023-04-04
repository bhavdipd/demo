package domain

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AlbumRepositoryDb struct {
	client *mongo.Client
}

func (d AlbumRepositoryDb) FindAll() ([]Album, error) {

	/*
	   List databases
	*/
	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(databases)

	coll := d.client.Database("studio").Collection("albums")

	albumCollection, err := coll.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	albums := make([]Album, 0)
	for albumCollection.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem Album
		err := albumCollection.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		albums = append(albums, elem)

	}

	return albums, nil
}

func NewAlbumRepositoryDb() AlbumRepositoryDb {
	/*
	   Connect to my cluster
	*/
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)

	return AlbumRepositoryDb{client}
}
