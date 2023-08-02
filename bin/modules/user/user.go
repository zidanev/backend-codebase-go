package user

import (
	"context"

	"codebase-go/bin/modules/user/models"
	"codebase-go/bin/pkg/utils"
)

type UsecaseQuery interface {
	// idiomatic go, ctx first before payload. See https://pkg.go.dev/context#pkg-overview
	GetUser(ctx context.Context, userId string) utils.Result
}

type UsecaseCommand interface {
	// idiomatic go, ctx first before payload. See https://pkg.go.dev/context#pkg-overview
	RegisterUser(ctx context.Context, payload models.User) utils.Result
	LoginUser(ctx context.Context, payload models.LoginRequest) utils.Result
	UpdateUser(ctx context.Context, payload models.User) utils.Result
	DeleteUser(ctx context.Context, userId string) utils.Result
}

type MongodbRepositoryQuery interface {
	// idiomatic go, ctx first before payload. See https://pkg.go.dev/context#pkg-overview
	FindOne(ctx context.Context, userId string) <-chan utils.Result
	FindOneByUsername(ctx context.Context, username string) <-chan utils.Result
}

type MongodbRepositoryCommand interface {
	// idiomatic go, ctx first before payload. See https://pkg.go.dev/context#pkg-overview
	NewObjectID(ctx context.Context) string
	InsertOneUser(ctx context.Context, data models.User) <-chan utils.Result
	UpdateOneUser(ctx context.Context, data models.User) <-chan utils.Result
	DeleteOneUser(ctx context.Context, userId string) <-chan utils.Result
}
