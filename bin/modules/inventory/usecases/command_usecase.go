package usecases

import (
	"codebase-go/bin/modules/inventory"
	"codebase-go/bin/modules/inventory/models"
	"codebase-go/bin/pkg/errors"
	"codebase-go/bin/pkg/redis"
	"codebase-go/bin/pkg/utils"
	"context"
)

type commandUsecase struct {
	inventoryRepositoryQuery   inventory.MongodbRepositoryQuery
	inventoryRepositoryCommand inventory.MongodbRepositoryCommand
	redis                 redis.Collections
}

func NewCommandUsecase(mq inventory.MongodbRepositoryQuery, mc inventory.MongodbRepositoryCommand, rc redis.Collections) inventory.UsecaseCommand {
	return &commandUsecase{
		inventoryRepositoryQuery:   mq,
		inventoryRepositoryCommand: mc,
		redis:                 rc,
	}
}

func (c commandUsecase) CreateInventory(ctx context.Context, payload models.Inventory) utils.Result {
	var result utils.Result = <-c.inventoryRepositoryCommand.InsertOneInventory(ctx, payload)
	if result.Error != nil {
		errObj := errors.InternalServerError("Failed insert inventory")
		result.Error = errObj
		return result
	}

	return result
}
