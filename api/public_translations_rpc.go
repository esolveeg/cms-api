package api

import (
	"context"

	"connectrpc.com/connect"
	"github.com/esolveeg/cms-api/proto_gen/devkit/v1"
)

func (api *Api) TranslationsDelete(ctx context.Context, req *connect.Request[devkitv1.TranslationsDeleteRequest]) (*connect.Response[devkitv1.TranslationsDeleteResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	response, err := api.publicUsecase.TranslationsDelete(ctx, req.Msg)
	return connect.NewResponse(response), err
}

func (api *Api) TranslationsList(ctx context.Context, req *connect.Request[devkitv1.TranslationsListRequest]) (*connect.Response[devkitv1.TranslationsListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	response, err := api.publicUsecase.TranslationsList(ctx)
	return connect.NewResponse(response), err
}

func (api *Api) TranslationsCreateUpdateBulk(ctx context.Context, req *connect.Request[devkitv1.TranslationsCreateUpdateBulkRequest]) (*connect.Response[devkitv1.TranslationsCreateUpdateBulkResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	resp, err := api.publicUsecase.TranslationsCreateUpdateBulk(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
