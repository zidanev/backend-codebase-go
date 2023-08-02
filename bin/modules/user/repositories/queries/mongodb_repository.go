package queries

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"codebase-go/bin/modules/user"
	"codebase-go/bin/modules/user/models"
	"codebase-go/bin/pkg/databases/mongodb"
	"codebase-go/bin/pkg/utils"
)

type queryMongodbRepository struct {
	mongoDb mongodb.MongoDBLogger
}

func NewQueryMongodbRepository(mongodb mongodb.MongoDBLogger) user.MongodbRepositoryQuery {
	return &queryMongodbRepository{
		mongoDb: mongodb,
	}
}

func (q queryMongodbRepository) FindOne(ctx context.Context, userId string) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		objId, _ := primitive.ObjectIDFromHex(userId)
		var user models.User
		err := q.mongoDb.FindOne(mongodb.FindOne{
			Result:         &user,
			CollectionName: "users",
			Filter: bson.M{
				"_id": objId,
			},
		}, ctx)
		if err != nil {
			output <- utils.Result{
				Error: err,
			}
		}

		output <- utils.Result{
			Data: user,
		}

	}()

	return output
}

func (q queryMongodbRepository) FindOneByUsername(ctx context.Context, username string) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		var user models.User
		err := q.mongoDb.FindOne(mongodb.FindOne{
			Result:         &user,
			CollectionName: "users",
			Filter: bson.M{
				"username": username,
			},
		}, ctx)

		if err != nil {
			output <- utils.Result{
				Error: err,
			}
		}

		output <- utils.Result{
			Data: user,
		}

	}()

	return output
}
