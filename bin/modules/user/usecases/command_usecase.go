package usecases

import (
	"codebase-go/bin/config"
	"codebase-go/bin/modules/user"
	"codebase-go/bin/modules/user/models"
	"codebase-go/bin/pkg/errors"
	"codebase-go/bin/pkg/redis"
	"codebase-go/bin/pkg/token"
	"codebase-go/bin/pkg/utils"
	"context"
)

type commandUsecase struct {
	userRepositoryQuery   user.MongodbRepositoryQuery
	userRepositoryCommand user.MongodbRepositoryCommand
	redis                 redis.Collections
}

func NewCommandUsecase(mq user.MongodbRepositoryQuery, mc user.MongodbRepositoryCommand, rc redis.Collections) user.UsecaseCommand {
	return &commandUsecase{
		userRepositoryQuery:   mq,
		userRepositoryCommand: mc,
		redis:                 rc,
	}
}

func (c commandUsecase) RegisterUser(ctx context.Context, payload models.User) utils.Result {
	var result utils.Result

	queryRes := <-c.userRepositoryQuery.FindOneByUsername(ctx, payload.Username)
	if queryRes.Data != nil {
		errObj := errors.Conflict("User already exist")
		result.Error = errObj
		return result
	}

	payload.Password = utils.HashPassword(payload.Password)

	result = <-c.userRepositoryCommand.InsertOneUser(ctx, payload)
	if result.Error != nil {
		errObj := errors.InternalServerError("Failed insert user")
		result.Error = errObj
		return result
	}

	return result
}

func (c commandUsecase) LoginUser(ctx context.Context, payload models.LoginRequest) utils.Result {
	var result utils.Result

	queryRes := <-c.userRepositoryQuery.FindOneByUsername(ctx, payload.Username)
	if queryRes.Data == nil {
		errObj := errors.NotFound("User not found")
		result.Error = errObj
		return result
	}

	user := queryRes.Data.(models.User)
	valid := utils.CheckPasswordHash(payload.Password, user.Password)
	if !valid {
		errObj := errors.UnauthorizedError("Password not match")
		result.Error = errObj
		return result
	}

	claim := token.Claim{
		Username: user.Username,
		UserId:   user.Id,
	}

	jwt := <-token.Generate(ctx, config.GetConfig().PrivateKey, &claim, config.GetConfig().AccessTokenExpired)
	if jwt.Error != nil {
		errObj := errors.BadRequest("Invalid token")
		result.Error = errObj
		return result
	}
	data := models.LoginResponse{
		Id:          user.Id,
		Username:    user.Username,
		Email:       user.Email,
		AccessToken: jwt.Data.(string),
	}
	result.Data = data
	return result
}

func (c commandUsecase) UpdateUser(ctx context.Context, payload models.User) utils.Result {
	var result utils.Result

	queryRes := <-c.userRepositoryQuery.FindOne(ctx, payload.Id)
	if queryRes.Data == nil {
		errObj := errors.NotFound("User not found")
		result.Error = errObj
		return result
	}

	payload.Password = utils.HashPassword(payload.Password)

	result = <-c.userRepositoryCommand.UpdateOneUser(ctx, payload)
	if result.Error != nil {
		errObj := errors.InternalServerError("Failed update user")
		result.Error = errObj
		return result
	}

	return result
}

func (c commandUsecase) DeleteUser(ctx context.Context, userId string) utils.Result {
	var result utils.Result

	queryRes := <-c.userRepositoryQuery.FindOne(ctx, userId)
	if queryRes.Data == nil {
		errObj := errors.NotFound("User not found")
		result.Error = errObj
		return result
	}

	result = <-c.userRepositoryCommand.DeleteOneUser(ctx, userId)
	if result.Error != nil {
		errObj := errors.InternalServerError("Failed delete user")
		result.Error = errObj
		return result
	}

	return result
}
