package usecase

import (
	"context"
	apiv1 "github.com/esolveeg/cms-api/proto_gen/devkit/v1"
	"github.com/google/uuid"
	"github.com/supabase-community/auth-go/types"
)

func (u *AccountsUsecase) UserDelete(ctx context.Context, userID int32) (*apiv1.AccountsSchemaUser, error) {
	user, err := u.repo.UserDelete(ctx, userID)
	if err != nil {
		return nil, err
	}
	err = u.redisClient.AuthSessionDelete(ctx, user.UserEmail)
	if err != nil {
		return nil, err
	}
	err = u.redisClient.AuthSessionDelete(ctx, user.UserName)
	if err != nil {
		return nil, err
	}
	err = u.redisClient.AuthSessionDelete(ctx, user.UserPhone.String)
	if err != nil {
		return nil, err
	}

	resp := u.adapter.UserEntityGrpcFromSql(user)
	return resp, nil
}

func (u *AccountsUsecase) UsersDeleteRestore(ctx context.Context, req *apiv1.UsersDeleteRestoreRequest) (*apiv1.UsersDeleteRestoreResponse, error) {
	err := u.repo.UsersDeleteRestore(ctx, req.Records)
	if err != nil {
		return nil, err
	}
	return &apiv1.UsersDeleteRestoreResponse{}, nil
}
func (u *AccountsUsecase) UsersList(ctx context.Context) (*apiv1.UsersListResponse, error) {
	users, err := u.repo.UsersList(ctx)
	if err != nil {
		return nil, err
	}
	response := u.adapter.UsersListGrpcFromSql(users)
	return response, nil
}
func (u *AccountsUsecase) UserCreateUpdate(ctx context.Context, req *apiv1.UserCreateUpdateRequest) (*apiv1.UserCreateUpdateResponse, error) {
	supabasRequest := types.AdminUpdateUserRequest{
		Email:    req.UserEmail,
		Password: req.UserPassword,
	}
	if req.UserId != 0 {
		userID, err := u.repo.AuthUserIDFindByEmail(ctx, req.UserEmail)
		if err != nil {
			return nil, err
		}
		uuid, err := uuid.Parse(*userID)
		if err != nil {
			return nil, err
		}
		supabasRequest.UserID = uuid
	}
	_, err := u.supaapi.UserCreateUpdate(supabasRequest)
	if err != nil {
		return nil, err
	}
	userCreateParams := u.adapter.UserCreateUpdateSqlFromGrpc(req)
	user, err := u.repo.UserCreateUpdate(ctx, *userCreateParams)
	if err != nil {
		return nil, err
	}

	resp := u.adapter.UserCreateUpdateGrpcFromSql(user)
	return resp, nil
}
