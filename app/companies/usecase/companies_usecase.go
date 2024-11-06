package usecase

import (
	"context"
	"github.com/esolveeg/cms-api/proto_gen/devkit/v1"
)

func (u *CompaniesUsecase) CompaniesCreateUpdate(ctx context.Context, req *devkitv1.CompaniesCreateUpdateRequest) (*devkitv1.CompaniesCreateUpdateResponse, error) {
	sqlReq := u.adapter.CompaniesCreateUpdateSqlFromGrpc(req)
	record, err := u.repo.CompaniesCreateUpdate(ctx, sqlReq)

	if err != nil {
		return nil, err
	}
	resp := u.adapter.CompanyEntityGrpcFromSql(record)
	return &devkitv1.CompaniesCreateUpdateResponse{Record: resp}, nil

}

func (u *CompaniesUsecase) CompaniesList(ctx context.Context, req *devkitv1.CompaniesListRequest) (*devkitv1.CompaniesListResponse, error) {
	record, err := u.repo.CompaniesList(ctx)
	if err != nil {
		return nil, err
	}
	resp := u.adapter.CompaniesListGrpcFromSql(&record)
	return resp, nil
}
