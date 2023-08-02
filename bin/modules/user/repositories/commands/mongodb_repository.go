package commands

import (
	"context"

	"codebase-go/bin/modules/user"
	"codebase-go/bin/modules/user/models"
	"codebase-go/bin/pkg/databases/mongodb"
	"codebase-go/bin/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type commandMongodbRepository struct {
	mongoDb mongodb.MongoDBLogger
}

func NewCommandMongodbRepository(mongodb mongodb.MongoDBLogger) user.MongodbRepositoryCommand {
	return &commandMongodbRepository{
		mongoDb: mongodb,
	}
}

func (c commandMongodbRepository) NewObjectID(ctx context.Context) string {
	return primitive.NewObjectID().Hex()
}

func (c commandMongodbRepository) InsertOneUser(ctx context.Context, data models.User) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)
		err := c.mongoDb.InsertOne(mongodb.InsertOne{
			CollectionName: "users",
			Document:       data,
		}, ctx)

		if err != nil {
			output <- utils.Result{Error: err}
		}
	}()

	return output
}

func (c commandMongodbRepository) UpdateOneUser(ctx context.Context, data models.User) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		objId, err := primitive.ObjectIDFromHex(data.Id)
		if err != nil {
			output <- utils.Result{Error: err}
		}

		err = c.mongoDb.UpdateOne(mongodb.UpdateOne{
			CollectionName: "users",
			Document:       data.UpsertUser(),
			Filter: bson.M{
				"_id": objId,
			},
		}, ctx)

		if err != nil {
			output <- utils.Result{Error: err}
		}

		output <- utils.Result{Data: data}
	}()

	return output
}

func (c commandMongodbRepository) DeleteOneUser(ctx context.Context, userId string) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		objId, err := primitive.ObjectIDFromHex(userId)
		if err != nil {
			output <- utils.Result{Error: err}
		}

		err = c.mongoDb.DeleteOne(mongodb.DeleteOne{
			CollectionName: "users",
			Filter: bson.M{
				"_id": objId,
			},
		}, ctx)

		if err != nil {
			output <- utils.Result{Error: err}
		}
	}()

	return output
}
