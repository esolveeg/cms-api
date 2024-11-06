package usecase

import (
	"bytes"
	"context"
	"io"

	"github.com/esolveeg/cms-api/proto_gen/devkit/v1"
	storage_go "github.com/supabase-community/storage-go"
)

func (s *PublicUsecase) UploadFile(ctx context.Context, req *devkitv1.UploadFileRequest) (*devkitv1.UploadFileResponse, error) {
	reader := io.NopCloser(bytes.NewReader(req.Reader))
	isUpsert := true
	fileOpts := storage_go.FileOptions{
		ContentType: &req.FileType,
		Upsert:      &isUpsert,
	}
	response, err := s.supaapi.StorageClient.UploadFile(req.BucketName, req.Path, reader, fileOpts)
	if err != nil {
		return nil, err
	}
	return &devkitv1.UploadFileResponse{
		Path: response.Key,
	}, nil

}
func (s *PublicUsecase) FilesList(ctx context.Context, req *devkitv1.FilesListRequest) (*devkitv1.FilesListResponse, error) {
	options := storage_go.FileSearchOptions{
		Limit:  int(req.Limit),
		Offset: int(req.Offest),
	}
	resp, err := s.supaapi.StorageClient.ListFiles(req.BucketId, req.QueryPath, options)
	if err != nil {
		return nil, err
	}

	response := s.adapter.FilesListGrpcFromSupa(resp)
	return response, nil
}
func (s *PublicUsecase) BucketCreateUpdate(ctx context.Context, req *devkitv1.BucketCreateUpdateRequest) (*devkitv1.BucketCreateUpdateResponse, error) {
	request := storage_go.BucketOptions{
		Public:           req.IsPulic,
		FileSizeLimit:    req.FileSizeLimit,
		AllowedMimeTypes: req.AllowedFileTypes,
	}
	if req.IsUpdate {
		_, err := s.supaapi.StorageClient.UpdateBucket(req.BucketName, request)
		if err != nil {
			return nil, err
		}
		return &devkitv1.BucketCreateUpdateResponse{}, nil

	}
	resp, err := s.supaapi.StorageClient.CreateBucket(req.BucketName, request)
	if err != nil {
		return nil, err
	}
	bucket := s.adapter.StorageBucketGrpcFromSupa(&resp)
	return &devkitv1.BucketCreateUpdateResponse{
		Bucket: bucket,
	}, nil
}

func (s *PublicUsecase) BucketsList(ctx context.Context, req *devkitv1.BucketsListRequest) (*devkitv1.BucketsListResponse, error) {
	resp, err := s.supaapi.StorageClient.ListBuckets()
	if err != nil {
		return nil, err
	}
	response := s.adapter.BucketsListGrpcFromSupa(resp)
	return response, nil
}

func (s *PublicUsecase) FilesDelete(ctx context.Context, req *devkitv1.FilesDeleteRequest) (*devkitv1.FilesDeleteResponse, error) {
	resp, err := s.supaapi.StorageClient.RemoveFile(req.BucketId, req.FilesPaths)
	if err != nil {
		return nil, err
	}
	response := s.adapter.FilesDeleteGrpcFromSupa(resp)
	return response, nil
}

func (s *PublicUsecase) UploadFiles(ctx context.Context, req *devkitv1.UploadFilesRequest) (*devkitv1.UploadFilesResponse, error) {
	images := make([]string, len(req.Files))
	for index, file := range req.Files {
		response, err := s.UploadFile(ctx, file)
		if err != nil {
			return nil, err
		}
		images[index] = response.Path
	}
	return &devkitv1.UploadFilesResponse{Path: images}, nil
}
