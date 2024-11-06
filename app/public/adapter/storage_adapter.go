package adapter

import (
	devkitv1 "github.com/esolveeg/cms-api/proto_gen/devkit/v1"
	storage_go "github.com/supabase-community/storage-go"
)

func (a *PublicAdapter) StorageBucketGrpcFromSupa(resp *storage_go.Bucket) *devkitv1.StorageBucket {
	return &devkitv1.StorageBucket{
		Name:      resp.Name,
		CreatedAt: resp.CreatedAt,
		Id:        resp.Id,
		Public:    resp.Public,
	}
}

func (a *PublicAdapter) FileUploadResponseGrpcFromSupa(resp *storage_go.FileUploadResponse) *devkitv1.FileUploadResponse {
	return &devkitv1.FileUploadResponse{
		Key:        resp.Key,
		Message:    resp.Message,
		StatusCode: resp.Code,
		Data:       resp.Data,
	}
}

func (a *PublicAdapter) FileObjectGrpcFromSupa(resp *storage_go.FileObject) *devkitv1.FileObject {
	return &devkitv1.FileObject{
		Name:      resp.Name,
		UpdatedAt: resp.UpdatedAt,
		BucketId:  resp.BucketId,
		CreatedAt: resp.CreatedAt,
		Id:        resp.Id,
	}
}

func (a *PublicAdapter) FilesDeleteGrpcFromSupa(resp []storage_go.FileUploadResponse) *devkitv1.FilesDeleteResponse {
	response := make([]*devkitv1.FileUploadResponse, len(resp))
	for index, rec := range resp {
		response[index] = a.FileUploadResponseGrpcFromSupa(&rec)
	}
	return &devkitv1.FilesDeleteResponse{
		Responses: response,
	}
}
func (a *PublicAdapter) FilesListGrpcFromSupa(resp []storage_go.FileObject) *devkitv1.FilesListResponse {
	files := make([]*devkitv1.FileObject, len(resp))
	for index, rec := range resp {
		files[index] = a.FileObjectGrpcFromSupa(&rec)
	}
	return &devkitv1.FilesListResponse{Files: files}
}
func (a *PublicAdapter) BucketsListGrpcFromSupa(resp []storage_go.Bucket) *devkitv1.BucketsListResponse {
	buckets := make([]*devkitv1.StorageBucket, len(resp))
	for index, rec := range resp {
		buckets[index] = a.StorageBucketGrpcFromSupa(&rec)
	}
	return &devkitv1.BucketsListResponse{Buckets: buckets}
}
