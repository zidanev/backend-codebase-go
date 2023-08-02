package inventory

import (
	"context"

	"codebase-go/bin/modules/inventory/models"
	"codebase-go/bin/pkg/utils"
)

type UsecaseQuery interface {
	// idiomatic go, ctx first before payload. See https://pkg.go.dev/context#pkg-overview
	GetInventory(ctx context.Context, userId string) utils.Result
}

type UsecaseCommand interface {
	// idiomatic go, ctx first before payload. See https://pkg.go.dev/context#pkg-overview
	CreateInventory(ctx context.Context, payload models.Inventory) utils.Result
}

type MongodbRepositoryQuery interface {
	// idiomatic go, ctx first before payload. See https://pkg.go.dev/context#pkg-overview
	FindOne(ctx context.Context, userId string) <-chan utils.Result
}

type MongodbRepositoryCommand interface {
	// idiomatic go, ctx first before payload. See https://pkg.go.dev/context#pkg-overview
	NewObjectID(ctx context.Context) string
	InsertOneInventory(ctx context.Context, data models.Inventory) <-chan utils.Result
}
