package usecase

import (
	"context"

	devkitv1 "github.com/esolveeg/cms-api/proto_gen/devkit/v1"
)

func (u *PublicUsecase) IconsCreateUpdateBulk(ctx context.Context, req *devkitv1.IconsCreateUpdateBulkRequest) (*devkitv1.IconsListResponse, error) {
	params := u.adapter.IconsCreateUpdateBulkSqlFromGrpc(req)
	icons, err := u.repo.IconsCreateUpdateBulk(ctx, params)
	if err != nil {
		return nil, err
	}
	res := u.adapter.IconsInputListGrpcFromSql(*icons)
	return res, nil
}
func (u *PublicUsecase) IconsInputList(ctx context.Context) (*devkitv1.IconsListResponse, error) {
	icons, err := u.repo.IconsInputList(ctx)
	if err != nil {
		return nil, err
	}
	res := u.adapter.IconsInputListGrpcFromSql(*icons)
	return res, nil
}
