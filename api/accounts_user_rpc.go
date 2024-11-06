package api

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	devkitv1 "github.com/esolveeg/cms-api/proto_gen/devkit/v1"
)

func (api *Api) UsersList(ctx context.Context, req *connect.Request[devkitv1.UsersListRequest]) (*connect.Response[devkitv1.UsersListResponse], error) {
	userPayload, err := api.authorizeRequestHeader(ctx, req.Header())
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("invalid access token: %w", err))
	}

	permissionMap, err := api.authorizedUserPermissions(ctx, userPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to get user permissions: %w", err)
	}

	if _, ok := permissionMap["users"]; !ok {
		return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("user does not have the required permissions"))
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

	response, err := api.accountsUsecase.UserCreateUpdate(ctx, req.Msg)
	if err != nil {
		return nil, fmt.Errorf("failed to create or update user: %w", err)
	}
	return connect.NewResponse(response), nil
}
func (api *Api) UserLogin(ctx context.Context, req *connect.Request[devkitv1.UserLoginRequest]) (*connect.Response[devkitv1.UserLoginResponse], error) {
	err := api.validator.Validate(req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	response, err := api.accountsUsecase.UserLogin(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}
func (api *Api) UserResetPassword(ctx context.Context, req *connect.Request[devkitv1.UserResetPasswordRequest]) (*connect.Response[devkitv1.UserResetPasswordResponse], error) {
	err := api.validator.Validate(req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("validation error: %w", err))
	}
	if req.Msg.NewPassword != req.Msg.NewPasswordConfirmation {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("password and confirmation do not match"))
	}
	response, err := api.accountsUsecase.UserResetPassword(ctx, req.Msg)
	if err != nil {
		return nil, fmt.Errorf("failed to reset password: %w", err)
	}
	return connect.NewResponse(response), nil
}

func (api *Api) UserAuthorize(ctx context.Context, req *connect.Request[devkitv1.UserAuthorizeRequest]) (*connect.Response[devkitv1.UserAuthorizeResponse], error) {
	payload, err := api.authorizeRequestHeader(ctx, req.Header())
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("invalid access token: %w", err))
	}
	response, _, err := api.accountsUsecase.AppLogin(ctx, payload.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to authorize user: %w", err)
	}
	return connect.NewResponse(&devkitv1.UserAuthorizeResponse{
		User:          response.User,
		NavigationBar: response.NavigationBar,
	}), nil
}
