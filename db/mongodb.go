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

func NewClient(dbConnection *DatabaseConnection) *Client {

	once.Do(func() {
		ctx := context.Background()
		clientOptions := options.Client().ApplyURI(dbConnection.Url)

		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		db := client.Database(dbConnection.Name)
		clientInstance = &Client{Client: client, Context: ctx, Database: db}
	})

	return clientInstance
}

func (c *Client) GetCollection(collName string) *mongo.Collection {
	c.Collection = c.Database.Collection(collName)

	return c.Collection
}

func (c *Client) Close() {

	if err := c.Client.Disconnect(c.Context); err != nil {
		log.Fatal(err)
	}
}
func (c *Client) DropDatabase() error {
	if err := c.Client.Database(c.Database.Name()).Drop(c.Context); err != nil {
		return fmt.Errorf("error dropping database: %w", err)
	}
	return nil
}

func (c *Client) DropCollection(collName string) error {

	if err := c.Database.Collection(collName).Drop(c.Context); err != nil {
		return fmt.Errorf("error dropping collection %s: %w", collName, err)
	}
	return nil
}
