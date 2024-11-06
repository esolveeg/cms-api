package api

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	devkitv1 "github.com/esolveeg/cms-api/proto_gen/devkit/v1"
)

func (api *Api) UsersList(ctx context.Context, req *connect.Request[devkitv1.UsersListRequest]) (*connect.Response[devkitv1.UsersListResponse], error) {
	_, err := api.checkForAccess(req.Header(), "users", "list")
	if err != nil {
		return nil, err
	}

	response, err := api.accountsUsecase.UsersList(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users list: %w", err)
	}
	return connect.NewResponse(response), nil
}

func (api *Api) UserCreateUpdate(ctx context.Context, req *connect.Request[devkitv1.UserCreateUpdateRequest]) (*connect.Response[devkitv1.UserCreateUpdateResponse], error) {
	err := api.validator.Validate(req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("validation error: %w", err))
	}
	if req.Msg.GetUserId() == 0 && (req.Msg.GetUserName() == "" || req.Msg.GetUserEmail() == "" || req.Msg.GetUserPhone() == "") {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("user ID is missing; name, email, and phone are required for creating a new user"))
	}
	_, err = api.checkForAccess(req.Header(), "users", "create_update")
	if err != nil {
		return nil, err
	}
	response, err := api.accountsUsecase.UserCreateUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}

func (api *Api) UsersDeleteRestore(ctx context.Context, req *connect.Request[devkitv1.UsersDeleteRestoreRequest]) (*connect.Response[devkitv1.UsersDeleteRestoreResponse], error) {
	_, err := api.checkForAccess(req.Header(), "users", "delete_restore")
	if err != nil {
		return nil, err
	}
	_, err = api.accountsUsecase.UsersDeleteRestore(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&devkitv1.UsersDeleteRestoreResponse{}), nil
}

func (api *Api) UserDelete(ctx context.Context, req *connect.Request[devkitv1.UserDeleteRequest]) (*connect.Response[devkitv1.UserDeleteResponse], error) {
	_, err := api.checkForAccess(req.Header(), "users", "delete")
	if req.Msg.RecordId != 0 {
		_, err := api.checkForAccess(req.Header(), "users", "delete_restore")
		if err != nil {
			return nil, err
		}
	}
	if req.Msg.RecordId == 0 {
		payload, err := api.authorizeRequestHeader(req.Header())
		if err != nil {
			return nil, err
		}
		req.Msg.RecordId = payload.UserId
	}
	_, err = api.accountsUsecase.UserDelete(ctx, req.Msg.RecordId)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&devkitv1.UserDeleteResponse{}), nil
}
