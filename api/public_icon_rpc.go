package api

import (
	"context"

	"connectrpc.com/connect"
	devkitv1 "github.com/esolveeg/cms-api/proto_gen/devkit/v1"
)

func (api *Api) IconsCreateUpdateBulk(ctx context.Context, req *connect.Request[devkitv1.IconsCreateUpdateBulkRequest]) (*connect.Response[devkitv1.IconsCreateUpdateBulkResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	_, err := api.checkForAccess(req.Header(), "icons", "create_update")
	if err != nil {
		return nil, err
	}
	_, err = api.publicUsecase.IconsCreateUpdateBulk(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&devkitv1.IconsCreateUpdateBulkResponse{}), nil
}

func (api *Api) IconsList(ctx context.Context, req *connect.Request[devkitv1.IconsListRequest]) (*connect.Response[devkitv1.IconsListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.IconsInputList(ctx)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
