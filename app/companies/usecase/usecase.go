package usecase

import (
	"context"
	"github.com/esolveeg/cms-api/app/companies/adapter"
	"github.com/esolveeg/cms-api/app/companies/repo"
	"github.com/esolveeg/cms-api/db"
	"github.com/esolveeg/cms-api/proto_gen/devkit/v1"
)

type CompaniesUsecaseInterface interface {
	// INJECT INTERFACE
	CompaniesList(ctx context.Context, req *devkitv1.CompaniesListRequest) (*devkitv1.CompaniesListResponse, error)
	CompaniesCreateUpdate(ctx context.Context, req *devkitv1.CompaniesCreateUpdateRequest) (*devkitv1.CompaniesCreateUpdateResponse, error)
}

type CompaniesUsecase struct {
	store   db.Store
	adapter adapter.CompaniesAdapterInterface
	repo    repo.CompaniesRepoInterface
}

func NewCompaniesUsecase(store db.Store) CompaniesUsecaseInterface {
	return &CompaniesUsecase{
		store:   store,
		adapter: adapter.NewCompaniesAdapter(),
		repo:    repo.NewCompaniesRepo(store),
	}
}
