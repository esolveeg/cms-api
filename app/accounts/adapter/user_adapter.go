package adapter

import (
	"encoding/json"

	"github.com/esolveeg/cms-api/db"
	devkitv1 "github.com/esolveeg/cms-api/proto_gen/devkit/v1"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

func (a *AccountsAdapter) UserFindNavigationBarsGrpcFromSql(resp *[]db.UserFindNavigationBarsRow) (*[]*devkitv1.NavigationBarItem, error) {
	response := make([]*devkitv1.NavigationBarItem, 0)
	for _, row := range *resp {
		items := make([]*devkitv1.NavigationBarItem, len(row.Items))
		if len(row.Items) == 0 && !row.Route.Valid {
			continue
		}
		if len(row.Items) > 0 {
			if err := json.Unmarshal(row.Items, &items); err != nil {
				return nil, err
			}
		}

		if row.LabelAr.Valid {
			row.LabelAr = pgtype.Text{Valid: true, String: row.LabelAr.String}
		} else {
			row.LabelAr = pgtype.Text{Valid: true, String: row.Label}
		}
		response = append(response, &devkitv1.NavigationBarItem{
			Key:     row.Key,
			Label:   row.Label,
			LabelAr: row.LabelAr.String,
			Icon:    row.Icon.String,
			Route:   row.Route.String,
			Items:   items,
		})
	}

	return &response, nil
}

func (a *AccountsAdapter) UserEntityGrpcFromSql(resp *db.AccountsSchemaUser) *devkitv1.AccountsSchemaUser {
	return &devkitv1.AccountsSchemaUser{
		UserId:            int32(resp.UserID),
		UserName:          resp.UserName,
		UserSecurityLevel: resp.UserSecurityLevel, // Security level of the user
		UserTypeId:        resp.UserTypeID,
		UserPhone:         resp.UserPhone.String,
		UserEmail:         resp.UserEmail, // User's email, unique in DB
		CreatedAt:         db.TimeToString(resp.CreatedAt.Time),
		DeletedAt:         db.TimeToString(resp.DeletedAt.Time),
	}
}

func (a *AccountsAdapter) UserCreateUpdateSqlFromGrpc(req *devkitv1.UserCreateUpdateRequest) *db.UserCreateUpdateParams {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.UserPassword), bcrypt.DefaultCost)
	resp := &db.UserCreateUpdateParams{
		UserID:            req.UserId,
		UserName:          req.UserName,
		UserSecurityLevel: req.UserSecurityLevel,
		UserTypeID:        req.UserTypeId,
		UserPhone:         req.UserPhone,
		UserEmail:         req.UserEmail,
		UserPassword:      string(hashedPassword),
		Roles:             req.Roles,
	}
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
