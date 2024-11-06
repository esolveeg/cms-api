package api

import (
	// INJECT IMPORTS
	"connectrpc.com/connect"
	"context"
	"github.com/esolveeg/cms-api/proto_gen/devkit/v1"
)

func (api *Api) CompaniesCreateUpdate(ctx context.Context, req *connect.Request[devkitv1.CompaniesCreateUpdateRequest]) (*connect.Response[devkitv1.CompaniesCreateUpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	err := api.validator.Validate(req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	resp, err := api.companiesUsecase.CompaniesCreateUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil

}
func (api *Api) CompaniesList(ctx context.Context, req *connect.Request[devkitv1.CompaniesListRequest]) (*connect.Response[devkitv1.CompaniesListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	err := api.validator.Validate(req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	resp, err := api.companiesUsecase.CompaniesList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil

}
