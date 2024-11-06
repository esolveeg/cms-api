package repo

import (
	// INJECT IMPORTS
	"context"
	"github.com/esolveeg/cms-api/db"
)

type CompaniesRepoInterface interface {
	// INJECT INTERFACE
	CompaniesList(ctx context.Context) ([]db.CompaniesSchemaCompany, error)
	CompaniesCreateUpdate(ctx context.Context, req *db.CompaniesCreateUpdateParams) (*db.CompaniesSchemaCompany, error)
}

type CompaniesRepo struct {
	store        db.Store
	errorHandler map[string]string
}

func NewCompaniesRepo(store db.Store) CompaniesRepoInterface {
	errorHandler := map[string]string{
		// INJECT ERRORS
		"companies_company_name_key":  "companyName",
		"companies_company_phone_key": "companyPhone",
		"companies_company_email_key": "companyEmail",
	}
	return &CompaniesRepo{
		store:        store,
		errorHandler: errorHandler,
	}
}
