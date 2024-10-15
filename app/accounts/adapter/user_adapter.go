package adapter

import (
	"github.com/darwishdev/devkit-api/db"
	devkitv1 "github.com/darwishdev/devkit-api/proto_gen/devkit/v1"
	"github.com/rs/zerolog/log"
)

func (a *AccountsAdapter) UserEntityGrpcFromSql(resp *db.AccountsSchemaUser) *devkitv1.AccountsSchemaUser {
	return &devkitv1.AccountsSchemaUser{
		UserId:            int32(resp.UserID),
		UserName:          resp.UserName,
		UserSecurityLevel: resp.UserSecurityLevel, // Security level of the user
		UserTypeId:        resp.UserTypeID,
		UserPhone:         resp.UserPhone.String,
		UserEmail:         resp.UserEmail,           // User's email, unique in DB
		UserPassword:      resp.UserPassword.String, // Password
		CreatedAt:         db.TimeToString(resp.CreatedAt.Time),
		DeletedAt:         db.TimeToString(resp.DeletedAt.Time),
	}
}

func (a *AccountsAdapter) UserCreateUpdateSqlFromGrpc(req *devkitv1.UserCreateUpdateRequest) *db.UserCreateUpdateParams {
	resp := &db.UserCreateUpdateParams{
		UserID:            req.UserId,
		UserName:          req.UserName,
		UserSecurityLevel: req.UserSecurityLevel,
		UserTypeID:        req.UserTypeId,
		UserPhone:         req.UserPhone,
		UserEmail:         req.UserEmail,
		UserPassword:      req.UserPassword,
		Roles:             req.Roles,
	}

	log.Debug().Interface("adapter here", resp).Msg("test adapter")
	return resp
}
func (a *AccountsAdapter) UsersListGrpcFromSql(resp []db.AccountsSchemaUser) *devkitv1.UsersListResponse {
	records := make([]*devkitv1.AccountsSchemaUser, 0)
	deletedRecords := make([]*devkitv1.AccountsSchemaUser, 0)
	for _, v := range resp {
		record := a.UserEntityGrpcFromSql(&v)
		if v.DeletedAt.Valid {
			deletedRecords = append(deletedRecords, record)
		} else {
			records = append(records, record)
		}
	}
	return &devkitv1.UsersListResponse{
		DeletedRecords: deletedRecords,
		Records:        records,
	}
}
func (a *AccountsAdapter) UserCreateUpdateGrpcFromSql(resp *db.AccountsSchemaUser) *devkitv1.UserCreateUpdateResponse {
	return &devkitv1.UserCreateUpdateResponse{
		User: a.UserEntityGrpcFromSql(resp),
	}
}