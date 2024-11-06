package repo

import (
	"context"

	"github.com/esolveeg/cms-api/db"
)

func (repo *PublicRepo) IconsCreateUpdateBulk(ctx context.Context, req db.IconsCreateUpdateBulkParams) (*[]db.Icon, error) {
	resp, err := repo.store.IconsCreateUpdateBulk(context.Background(), req)

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return &resp, nil
}
func (repo *PublicRepo) IconsInputList(ctx context.Context) (*[]db.Icon, error) {
	resp, err := repo.store.IconsInputList(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return &resp, nil
}
