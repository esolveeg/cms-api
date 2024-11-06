package usecase

import (
	"context"

	"github.com/esolveeg/cms-api/app/public/adapter"
	"github.com/esolveeg/cms-api/app/public/repo"
	"github.com/esolveeg/cms-api/db"
	"github.com/esolveeg/cms-api/pkg/redisclient"
	"github.com/esolveeg/cms-api/pkg/resend"
	devkitv1 "github.com/esolveeg/cms-api/proto_gen/devkit/v1"
	supaapigo "github.com/darwishdev/supaapi-go"
)

type PublicUsecaseInterface interface {
	TranslationsList(ctx context.Context) (*devkitv1.TranslationsListResponse, error)
	TranslationsCreateUpdateBulk(ctx context.Context, req *devkitv1.TranslationsCreateUpdateBulkRequest) (*devkitv1.TranslationsCreateUpdateBulkResponse, error)
	TranslationsDelete(ctx context.Context, req *devkitv1.TranslationsDeleteRequest) (*devkitv1.TranslationsDeleteResponse, error)
	FilesDelete(ctx context.Context, req *devkitv1.FilesDeleteRequest) (*devkitv1.FilesDeleteResponse, error)
	FilesList(ctx context.Context, req *devkitv1.FilesListRequest) (*devkitv1.FilesListResponse, error)
	SendEmail(ctx context.Context, req *devkitv1.SendEmailRequest) (*devkitv1.SendEmailResponse, error)
	BucketsList(ctx context.Context, req *devkitv1.BucketsListRequest) (*devkitv1.BucketsListResponse, error)
	SettingsUpdate(ctx context.Context, req *devkitv1.SettingsUpdateRequest) error
	SettingsFindForUpdate(ctx context.Context, req *devkitv1.SettingsFindForUpdateRequest) (*devkitv1.SettingsFindForUpdateResponse, error)
	UploadFile(ctx context.Context, req *devkitv1.UploadFileRequest) (*devkitv1.UploadFileResponse, error)
	BucketCreateUpdate(ctx context.Context, req *devkitv1.BucketCreateUpdateRequest) (*devkitv1.BucketCreateUpdateResponse, error)
	IconsCreateUpdateBulk(ctx context.Context, req *devkitv1.IconsCreateUpdateBulkRequest) (*devkitv1.IconsListResponse, error)
	IconsInputList(ctx context.Context) (*devkitv1.IconsListResponse, error)
	UploadFiles(ctx context.Context, req *devkitv1.UploadFilesRequest) (*devkitv1.UploadFilesResponse, error)
}

type PublicUsecase struct {
	store        db.Store
	repo         repo.PublicRepoInterface
	adapter      adapter.PublicAdapterInterface
	supaapi      supaapigo.Supaapi
	resendClient resend.ResendServiceInterface
	redisClient  redisclient.RedisClientInterface
}

func NewPublicUsecase(store db.Store, supaapi supaapigo.Supaapi, redisClient redisclient.RedisClientInterface, resendClient resend.ResendServiceInterface) PublicUsecaseInterface {
	return &PublicUsecase{
		resendClient: resendClient,
		supaapi:      supaapi,
		redisClient:  redisClient,
		adapter:      adapter.NewPublicAdapter(),
		repo:         repo.NewPublicRepo(store),
		store:        store,
	}
}
