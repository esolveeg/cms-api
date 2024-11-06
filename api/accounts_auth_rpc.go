package api

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	devkitv1 "github.com/darwishdev/devkit-api/proto_gen/devkit/v1"
)

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
	payload, err := api.authorizeRequestHeader(req.Header())
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("invalid access token: %w", err))
	}
	req.Msg.Email = payload.Username
	response, err := api.accountsUsecase.UserResetPassword(ctx, req.Msg)
	if err != nil {
		return nil, fmt.Errorf("failed to reset password: %w", err)
	}
	return connect.NewResponse(response), nil
}

func (api *Api) UserAuthorize(ctx context.Context, req *connect.Request[devkitv1.UserAuthorizeRequest]) (*connect.Response[devkitv1.UserAuthorizeResponse], error) {
	payload, err := api.authorizeRequestHeader(req.Header())
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
