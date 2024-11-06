package repo

import (
	"context"

	"github.com/esolveeg/cms-api/db"
)

type PublicRepoInterface interface {
	IconsCreateUpdateBulk(ctx context.Context, req db.IconsCreateUpdateBulkParams) (*[]db.Icon, error)
	TranslationsDelete(ctx context.Context, req []string) ([]db.Translation, error)
	TranslationsList(ctx context.Context) ([]db.Translation, error)
	TranslationsCreateUpdateBulk(ctx context.Context, req db.TranslationsCreateUpdateBulkParams) ([]db.TranslationsCreateUpdateBulkRow, error)
	SettingsUpdate(ctx context.Context, req *db.SettingsUpdateParams) error
	IconsInputList(ctx context.Context) (*[]db.Icon, error)
	SettingsFindForUpdate(ctx context.Context) (*[]db.SettingsFindForUpdateRow, error)
}

type PublicRepo struct {
	store        db.Store
	errorHandler map[string]string
}

func NewPublicRepo(store db.Store) PublicRepoInterface {
	errorHandler := map[string]string{
		"settings_setting_key_key": "roleName",
		"icons_icon_name_key":      "userName",
	}
	return &PublicRepo{
		store:        store,
		errorHandler: errorHandler,
	}
}
