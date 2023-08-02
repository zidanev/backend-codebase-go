package commands

import (
	"context"

	"codebase-go/bin/modules/inventory"
	"codebase-go/bin/modules/inventory/models"
	"codebase-go/bin/pkg/databases/mongodb"
	"codebase-go/bin/pkg/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type commandMongodbRepository struct {
	mongoDb mongodb.MongoDBLogger
}

func NewCommandMongodbRepository(mongodb mongodb.MongoDBLogger) inventory.MongodbRepositoryCommand {
	return &commandMongodbRepository{
		mongoDb: mongodb,
	}
}

func (c commandMongodbRepository) NewObjectID(ctx context.Context) string {
	return primitive.NewObjectID().Hex()
}

func (c commandMongodbRepository) InsertOneInventory(ctx context.Context, data models.Inventory) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)
		err := c.mongoDb.InsertOne(mongodb.InsertOne{
			CollectionName: "inventories",
			Document:       data,
		}, ctx)

		if err != nil {
			output <- utils.Result{Error: err}
		}
	}()

	return output
}