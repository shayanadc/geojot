package container

import (
	"geo-jot/db"
	"sync"
)

type ServiceContainer struct {
	dbClient *db.Client
}

var (
	container *ServiceContainer
	once      sync.Once
)

func GetContainer() *ServiceContainer {
	once.Do(func() {
		container = &ServiceContainer{}
	})
	return container
}

func (c *ServiceContainer) SetDBClient(client *db.Client) {
	c.dbClient = client
}

func (c *ServiceContainer) GetDBClient() *db.Client {
	return c.dbClient
}
