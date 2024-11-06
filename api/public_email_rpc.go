package api

import (
	"connectrpc.com/connect"
	"context"
	"github.com/esolveeg/cms-api/proto_gen/devkit/v1"
)

func (api *Api) SendEmail(ctx context.Context, req *connect.Request[devkitv1.SendEmailRequest]) (*connect.Response[devkitv1.SendEmailResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	response, err := api.publicUsecase.SendEmail(ctx, req.Msg)
	return connect.NewResponse(response), err
}
