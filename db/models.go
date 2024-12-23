// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type AccountsSchemaNavigationBar struct {
	NavigationBarID   int32  `json:"navigation_bar_id"`
	NavigationBarName string `json:"navigation_bar_name"`
}

type AccountsSchemaNavigationBarItem struct {
	NavigationBarItemID int32       `json:"navigation_bar_item_id"`
	MenuKey             string      `json:"menu_key"`
	Label               string      `json:"label"`
	LabelAr             pgtype.Text `json:"label_ar"`
	Icon                pgtype.Text `json:"icon"`
	Route               pgtype.Text `json:"route"`
	NavigationBarID     pgtype.Int4 `json:"navigation_bar_id"`
	ParentID            pgtype.Int4 `json:"parent_id"`
	PermissionID        pgtype.Int4 `json:"permission_id"`
}

type AccountsSchemaPermission struct {
	PermissionID          int32       `json:"permission_id"`
	PermissionFunction    string      `json:"permission_function"`
	PermissionName        string      `json:"permission_name"`
	PermissionDescription pgtype.Text `json:"permission_description"`
	PermissionGroup       string      `json:"permission_group"`
}

type AccountsSchemaRole struct {
	RoleID            int32            `json:"role_id"`
	RoleName          string           `json:"role_name"`
	RoleSecurityLevel int32            `json:"role_security_level"`
	RoleDescription   pgtype.Text      `json:"role_description"`
	CreatedAt         pgtype.Timestamp `json:"created_at"`
	UpdatedAt         pgtype.Timestamp `json:"updated_at"`
	DeletedAt         pgtype.Timestamp `json:"deleted_at"`
}

type AccountsSchemaRolePermission struct {
	RoleID       int32 `json:"role_id"`
	PermissionID int32 `json:"permission_id"`
}

type AccountsSchemaUser struct {
	UserID       int32            `json:"user_id"`
	UserName     string           `json:"user_name"`
	UserTypeID   int32            `json:"user_type_id"`
	UserPhone    pgtype.Text      `json:"user_phone"`
	UserEmail    string           `json:"user_email"`
	UserPassword pgtype.Text      `json:"user_password"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
	DeletedAt    pgtype.Timestamp `json:"deleted_at"`
}

type AccountsSchemaUserRole struct {
	UserID int32 `json:"user_id"`
	RoleID int32 `json:"role_id"`
}

type AccountsSchemaUserType struct {
	UserTypeID   int32  `json:"user_type_id"`
	UserTypeName string `json:"user_type_name"`
}

type CompaniesSchemaCompany struct {
	CompanyID               int32            `json:"company_id"`
	CompanyName             string           `json:"company_name"`
	CompanyNameAr           pgtype.Text      `json:"company_name_ar"`
	CompanyPhone            pgtype.Text      `json:"company_phone"`
	CompanyAddress          pgtype.Text      `json:"company_address"`
	CompanyAddressAr        pgtype.Text      `json:"company_address_ar"`
	CompanyDescription      pgtype.Text      `json:"company_description"`
	CompanyDescriptionAr    pgtype.Text      `json:"company_description_ar"`
	CompanyEmail            pgtype.Text      `json:"company_email"`
	CompanyLogo             pgtype.Text      `json:"company_logo"`
	CompanyLogoVertical     pgtype.Text      `json:"company_logo_vertical"`
	CompanyLogoDark         pgtype.Text      `json:"company_logo_dark"`
	CompanyLogoDarkVertical pgtype.Text      `json:"company_logo_dark_vertical"`
	CompanyValues           pgtype.Text      `json:"company_values"`
	CompanyMission          pgtype.Text      `json:"company_mission"`
	CompanyVision           pgtype.Text      `json:"company_vision"`
	CreatedAt               pgtype.Timestamp `json:"created_at"`
	UpdatedAt               pgtype.Timestamp `json:"updated_at"`
	DeletedAt               pgtype.Timestamp `json:"deleted_at"`
}

type Icon struct {
	IconID      int32  `json:"icon_id"`
	IconName    string `json:"icon_name"`
	IconContent string `json:"icon_content"`
}

type InputType struct {
	InputTypeID   int32  `json:"input_type_id"`
	InputTypeName string `json:"input_type_name"`
}

type Log struct {
	LogID          int32            `json:"log_id"`
	LogTitle       string           `json:"log_title"`
	UserID         int32            `json:"user_id"`
	RecordID       pgtype.Int4      `json:"record_id"`
	ActionName     pgtype.Text      `json:"action_name"`
	TableName      pgtype.Text      `json:"table_name"`
	PermissionName pgtype.Text      `json:"permission_name"`
	CreatedAt      pgtype.Timestamp `json:"created_at"`
}

type Setting struct {
	SettingID    int32            `json:"setting_id"`
	InputTypeID  int32            `json:"input_type_id"`
	SettingKey   string           `json:"setting_key"`
	SettingValue string           `json:"setting_value"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
}

type Tag struct {
	Tag string `json:"tag"`
}

type Translation struct {
	TranslationKey string `json:"translation_key"`
	ArabicValue    string `json:"arabic_value"`
	EnglishValue   string `json:"english_value"`
}
