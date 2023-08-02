package usecases

import (
	"context"

	"codebase-go/bin/modules/user"
	"codebase-go/bin/modules/user/models"
	"codebase-go/bin/pkg/errors"
	"codebase-go/bin/pkg/redis"
	"codebase-go/bin/pkg/utils"
)

type queryUsecase struct {
	userRepositoryQuery user.MongodbRepositoryQuery
}

func NewQueryUsecase(mq user.MongodbRepositoryQuery, rc redis.Collections) user.UsecaseQuery {
	return &queryUsecase{
		userRepositoryQuery: mq,
	}
}

func (q queryUsecase) GetUser(ctx context.Context, userId string) utils.Result {
	var result utils.Result

	queryRes := <-q.userRepositoryQuery.FindOne(ctx, userId)

	if queryRes.Error != nil {
		errObj := errors.InternalServerError("Internal server error")
		result.Error = errObj
		return result
	}
	user := queryRes.Data.(models.User)
	res := models.GetUserResponse{
		Id:           user.Id,
		Username:     user.Username,
		Email:        user.Email,
		FullName:     user.FullName,
		MobileNumber: user.MobileNumber,
		Status:       user.Status,
	}
	result.Data = res
	return result
}
