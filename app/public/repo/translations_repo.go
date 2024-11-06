package repo

import (
	"context"

	"github.com/esolveeg/cms-api/db"
)

func (repo *PublicRepo) TranslationsList(ctx context.Context) ([]db.Translation, error) {
	response, err := repo.store.TranslationsList(ctx)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return response, nil
}

func (repo *PublicRepo) TranslationsCreateUpdateBulk(ctx context.Context, req db.TranslationsCreateUpdateBulkParams) ([]db.TranslationsCreateUpdateBulkRow, error) {
	response, err := repo.store.TranslationsCreateUpdateBulk(ctx, req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return response, nil
}
func (repo *PublicRepo) TranslationsDelete(ctx context.Context, req []string) ([]db.Translation, error) {
	response, err := repo.store.TranslationsDelete(ctx, req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return response, nil
}
