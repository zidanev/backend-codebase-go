package queries

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"codebase-go/bin/modules/inventory"
	"codebase-go/bin/modules/inventory/models"
	"codebase-go/bin/pkg/databases/mongodb"
	"codebase-go/bin/pkg/utils"
)

type queryMongodbRepository struct {
	mongoDb mongodb.MongoDBLogger
}

func NewQueryMongodbRepository(mongodb mongodb.MongoDBLogger) inventory.MongodbRepositoryQuery {
	return &queryMongodbRepository{
		mongoDb: mongodb,
	}
}

func (q queryMongodbRepository) FindOne(ctx context.Context, kode string) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		var item models.Inventory
		err := q.mongoDb.FindOne(mongodb.FindOne{
			Result:         &item,
			CollectionName: "inventories",
			Filter: bson.M{
				"kode": kode,
			},
		}, ctx)
		if err != nil {
			output <- utils.Result{
				Error: err,
			}
		}

		output <- utils.Result{
			Data: item,
		}

	}()

	return output
}
