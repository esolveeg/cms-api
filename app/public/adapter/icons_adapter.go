package adapter

import (
	"github.com/esolveeg/cms-api/db"
	devkitv1 "github.com/esolveeg/cms-api/proto_gen/devkit/v1"
)

func (a *PublicAdapter) IconGrpcFromSql(icon *db.Icon) *devkitv1.Icon {
	return &devkitv1.Icon{
		IconId:      icon.IconID,
		IconName:    icon.IconName,
		IconContent: icon.IconContent,
	}

}
func (a *PublicAdapter) IconsCreateUpdateBulkSqlFromGrpc(req *devkitv1.IconsCreateUpdateBulkRequest) db.IconsCreateUpdateBulkParams {
	names := make([]string, len(req.Icons))
	contents := make([]string, len(req.Icons))
	for index, v := range req.Icons {
		names[index] = v.IconName
		contents[index] = v.IconContent
	}
	return db.IconsCreateUpdateBulkParams{
		IconsName:     names,
		IconsContents: contents,
	}
}
func (a *PublicAdapter) IconsInputListGrpcFromSql(resp []db.Icon) *devkitv1.IconsListResponse {
	records := make([]*devkitv1.Icon, len(resp))
	for index, v := range resp {
		records[index] = a.IconGrpcFromSql(&v)
	}
	return &devkitv1.IconsListResponse{
		Icons: records,
	}
}
