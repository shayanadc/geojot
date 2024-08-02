package db

import (
	"context"
	"fmt"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	Client     *mongo.Client
	Context    context.Context
	Collection *mongo.Collection
	Database   *mongo.Database
}

var (
	clientInstance *Client
	once           sync.Once
)

func NewClient() *Client {

	once.Do(func() {
		ctx := context.Background()
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")

		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		db := client.Database("geo-jot-db")
		clientInstance = &Client{Client: client, Context: ctx, Database: db}
	})

	return clientInstance
}

func (c *Client) GetCollection(collName string) *mongo.Collection {
	c.Collection = c.Database.Collection(collName)

	return c.Collection
}

func (c *Client) Close() {
	fmt.Println("DISCONNECTING")
	if err := c.Client.Disconnect(c.Context); err != nil {
		log.Fatal(err)
	}
}
