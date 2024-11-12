package usecase

import (
	"context"

	devkitv1 "github.com/darwishdev/devkit-api/proto_gen/devkit/v1"
)

func (u *PublicUsecase) IconCreateUpdateBulk(ctx context.Context, req *devkitv1.IconCreateUpdateBulkRequest) (*devkitv1.IconListResponse, error) {
	params := u.adapter.IconCreateUpdateBulkSqlFromGrpc(req)
	icons, err := u.repo.IconCreateUpdateBulk(ctx, params)
	if err != nil {
		return nil, err
	}
	res := u.adapter.IconListGrpcFromSql(*icons)
	return res, nil
}
func (u *PublicUsecase) IconList(ctx context.Context) (*devkitv1.IconListResponse, error) {
	icons, err := u.repo.IconList(ctx)
	if err != nil {
		return nil, err
	}
	res := u.adapter.IconListGrpcFromSql(*icons)
	return res, nil
}
