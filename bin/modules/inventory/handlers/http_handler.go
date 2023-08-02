package handlers

import (
	"codebase-go/bin/middlewares"
	"codebase-go/bin/modules/inventory"
	"codebase-go/bin/modules/inventory/models"
	"codebase-go/bin/pkg/errors"
	"codebase-go/bin/pkg/helpers"
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo"
)

type inventoryHttpHandler struct {
	inventoryUsecaseQuery   inventory.UsecaseQuery
	inventoryUseCaseCommand inventory.UsecaseCommand
}

func InitinventoryHttpHandler(e *echo.Echo, uq inventory.UsecaseQuery, uc inventory.UsecaseCommand) {

	handler := &inventoryHttpHandler{
		inventoryUsecaseQuery:   uq,
		inventoryUseCaseCommand: uc,
	}

	route := e.Group("/codebase-go")

	route.GET("/inventory/v1", handler.GetInventory, middlewares.VerifyBearer)
	route.POST("/inventory/v1/create", handler.CreateInventory, middlewares.VerifyBearer)
}

func (u inventoryHttpHandler) GetInventory(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	type ItemRequest struct {
		Kode string `json:"code"`
	}
	var itemReq ItemRequest
	err = json.Unmarshal(body, &itemReq)
	if err != nil {
		return err
	}

	var kode = itemReq.Kode

	result := u.inventoryUsecaseQuery.GetInventory(c.Request().Context(), kode)

	if result.Error != nil {
		return helpers.RespError(c, result.Error)
	}

	return helpers.RespSuccess(c, result.Data, "Get inventory success")
}

func (u inventoryHttpHandler) CreateInventory(c echo.Context) error {
	req := new(models.Inventory)


	if err := c.Bind(req); err != nil {
		return helpers.RespError(c, errors.BadRequest("bad request."))
	}
	if err := c.Validate(req); err != nil {
		return helpers.RespError(c, err)
	}

	result := u.inventoryUseCaseCommand.CreateInventory(c.Request().Context(), *req)
	if result.Error != nil {
		return helpers.RespError(c, result.Error)
	}

	return helpers.RespSuccess(c, result.Data, "Inventory Inserted successfully!")
}
