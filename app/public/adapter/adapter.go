package adapter

import (
	"github.com/esolveeg/cms-api/db"
	devkitv1 "github.com/esolveeg/cms-api/proto_gen/devkit/v1"
	"github.com/resend/resend-go/v2"
	storage_go "github.com/supabase-community/storage-go"
)

type PublicAdapterInterface interface {
	IconsCreateUpdateBulkSqlFromGrpc(req *devkitv1.IconsCreateUpdateBulkRequest) db.IconsCreateUpdateBulkParams
	SendEmailResendFromGrpc(req *devkitv1.SendEmailRequest) resend.SendEmailRequest
	TranslationsCreateUpdateBulkGrpcFromSql(resp []db.TranslationsCreateUpdateBulkRow) devkitv1.TranslationsListResponse
	TranslationsListGrpcFromSql(resp []db.Translation) devkitv1.TranslationsListResponse
	TranslationGrpcFromSql(resp *db.Translation) *devkitv1.Translation
	TranslationsCreateUpdateBulkSqlFromGrpc(req *devkitv1.TranslationsCreateUpdateBulkRequest) *db.TranslationsCreateUpdateBulkParams
	FilesDeleteGrpcFromSupa(resp []storage_go.FileUploadResponse) *devkitv1.FilesDeleteResponse
	FilesListGrpcFromSupa(resp []storage_go.FileObject) *devkitv1.FilesListResponse
	FileObjectGrpcFromSupa(resp *storage_go.FileObject) *devkitv1.FileObject
	FileUploadResponseGrpcFromSupa(resp *storage_go.FileUploadResponse) *devkitv1.FileUploadResponse
	BucketsListGrpcFromSupa(resp []storage_go.Bucket) *devkitv1.BucketsListResponse
	StorageBucketGrpcFromSupa(resp *storage_go.Bucket) *devkitv1.StorageBucket

	SettingsUpdateSqlFromGrpc(req *devkitv1.SettingsUpdateRequest) *db.SettingsUpdateParams
	SettingsEntityGrpcFromSql(resp []db.Setting) []*devkitv1.Setting
	SettingsFindForUpdateGrpcFromSql(resp *[]db.SettingsFindForUpdateRow) *devkitv1.SettingsFindForUpdateResponse
	IconsInputListGrpcFromSql(resp []db.Icon) *devkitv1.IconsListResponse
}

type PublicAdapter struct {
}

func NewPublicAdapter() PublicAdapterInterface {
	return &PublicAdapter{}
}
