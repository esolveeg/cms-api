package repo

import (
	"context"
	"github.com/esolveeg/cms-api/db"
)

func (repo *CompaniesRepo) CompaniesCreateUpdate(ctx context.Context, req *db.CompaniesCreateUpdateParams) (*db.CompaniesSchemaCompany, error) {
	resp, err := repo.store.CompaniesCreateUpdate(ctx, *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return &resp, nil
}
func (repo *CompaniesRepo) CompaniesList(ctx context.Context) ([]db.CompaniesSchemaCompany, error) {
	resp, err := repo.store.CompaniesList(ctx)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return resp, nil
}
