package adapter

import (
	"github.com/esolveeg/cms-api/db"
	"github.com/esolveeg/cms-api/proto_gen/devkit/v1"
)

func (a *CompaniesAdapter) CompaniesCreateUpdateSqlFromGrpc(req *devkitv1.CompaniesCreateUpdateRequest) *db.CompaniesCreateUpdateParams {
	return &db.CompaniesCreateUpdateParams{
		CompanyID:               req.CompanyId,
		CompanyName:             req.CompanyName,
		CompanyNameAr:           req.CompanyNameAr,
		CompanyPhone:            req.CompanyPhone,
		CompanyAddress:          req.CompanyAddress,
		CompanyAddressAr:        req.CompanyAddressAr,
		CompanyEmail:            req.CompanyEmail,
		CompanyDescription:      req.CompanyDescription,
		CompanyDescriptionAr:    req.CompanyDescriptionAr,
		CompanyLogo:             req.CompanyLogo,
		CompanyLogoVertical:     req.CompanyLogoVertical,
		CompanyLogoDark:         req.CompanyLogoDark,
		CompanyLogoDarkVertical: req.CompanyLogoDarkVertical,
	}
}

func (a *CompaniesAdapter) CompanyEntityGrpcFromSql(req *db.CompaniesSchemaCompany) *devkitv1.CompaniesSchemaCompany {
	return &devkitv1.CompaniesSchemaCompany{
		CompanyId:               req.CompanyID,
		CompanyName:             req.CompanyName,
		CompanyNameAr:           req.CompanyNameAr.String,
		CompanyPhone:            req.CompanyPhone.String,
		CompanyAddress:          req.CompanyAddress.String,
		CompanyAddressAr:        req.CompanyAddressAr.String,
		CompanyDescription:      req.CompanyDescription.String,
		CompanyDescriptionAr:    req.CompanyDescriptionAr.String,
		CompanyEmail:            req.CompanyEmail.String,
		CompanyLogo:             req.CompanyLogo.String,
		CompanyLogoVertical:     req.CompanyLogoVertical.String,
		CompanyLogoDark:         req.CompanyLogoDark.String,
		CompanyLogoDarkVertical: req.CompanyLogoDarkVertical.String,
		CreatedAt:               db.PgTimeToTimestamp(req.CreatedAt),
		UpdatedAt:               db.PgTimeToTimestamp(req.UpdatedAt),
		DeletedAt:               db.PgTimeToTimestamp(req.DeletedAt),
	}
}

func (a *CompaniesAdapter) CompaniesListGrpcFromSql(resp *[]db.CompaniesSchemaCompany) *devkitv1.CompaniesListResponse {
	records := make([]*devkitv1.CompaniesSchemaCompany, 0)
	deletedRecords := make([]*devkitv1.CompaniesSchemaCompany, 0)
	for _, v := range *resp {
		record := a.CompanyEntityGrpcFromSql(&v)
		if v.DeletedAt.Valid {
			deletedRecords = append(deletedRecords, record)
		} else {
			records = append(records, record)
		}
	}
	return &devkitv1.CompaniesListResponse{
		Records: records,
	}

}
