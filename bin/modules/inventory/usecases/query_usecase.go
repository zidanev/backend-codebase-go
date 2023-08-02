package usecases

import (
	"context"

	"codebase-go/bin/modules/inventory"
	"codebase-go/bin/modules/inventory/models"
	"codebase-go/bin/pkg/errors"
	"codebase-go/bin/pkg/redis"
	"codebase-go/bin/pkg/utils"
)

type queryUsecase struct {
	inventoryRepositoryQuery inventory.MongodbRepositoryQuery
}

func NewQueryUsecase(mq inventory.MongodbRepositoryQuery, rc redis.Collections) inventory.UsecaseQuery {
	return &queryUsecase{
		inventoryRepositoryQuery: mq,
	}
}

func (q queryUsecase) GetInventory(ctx context.Context, kode string) utils.Result {
	var result utils.Result

	queryRes := <-q.inventoryRepositoryQuery.FindOne(ctx, kode)

	if queryRes.Error != nil {
		errObj := errors.InternalServerError("Internal server error")
		result.Error = errObj
		return result
	}
	InvenGetInventory := queryRes.Data.(models.Inventory)
	res := models.Inventory{
		Id:           InvenGetInventory.Id,
		Nama:     		InvenGetInventory.Nama,
		Kode:        	InvenGetInventory.Kode,
		Harga:     		InvenGetInventory.Harga,
		Stock: 				InvenGetInventory.Stock,
		Kategori:     InvenGetInventory.Kategori,
	}
	result.Data = res
	return result
}
