package adapter

import (
	"github.com/esolveeg/cms-api/db"
	"github.com/esolveeg/cms-api/proto_gen/devkit/v1"
)

type CompaniesAdapterInterface interface {
	// INJECT INTERFACE
	CompaniesCreateUpdateSqlFromGrpc(req *devkitv1.CompaniesCreateUpdateRequest) *db.CompaniesCreateUpdateParams
	CompanyEntityGrpcFromSql(req *db.CompaniesSchemaCompany) *devkitv1.CompaniesSchemaCompany
	CompaniesListGrpcFromSql(resp *[]db.CompaniesSchemaCompany) *devkitv1.CompaniesListResponse
}

type CompaniesAdapter struct {
}

func NewCompaniesAdapter() CompaniesAdapterInterface {
	return &CompaniesAdapter{}
}
