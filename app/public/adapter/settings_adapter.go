package adapter

import (
	"github.com/esolveeg/cms-api/db"
	devkitv1 "github.com/esolveeg/cms-api/proto_gen/devkit/v1"
)

func (a *PublicAdapter) SettingsUpdateSqlFromGrpc(req *devkitv1.SettingsUpdateRequest) *db.SettingsUpdateParams {
	keys := make([]string, len(req.Settings))
	values := make([]string, len(req.Settings))
	for index, v := range req.Settings {
		keys[index] = v.SettingKey
		values[index] = v.SettingValue
	}
	return &db.SettingsUpdateParams{
		Keys:   keys,
		Values: values,
	}
}
func (a *PublicAdapter) SettingsEntityGrpcFromSql(resp []db.Setting) []*devkitv1.Setting {
	grpcResp := make([]*devkitv1.Setting, len(resp))
	for _, v := range resp {
		record := &devkitv1.Setting{
			SettingKey:   v.SettingKey,
			SettingValue: v.SettingValue,
		}
		grpcResp = append(grpcResp, record)
	}
	return grpcResp

}

func (a *PublicAdapter) SettingsFindForUpdateGrpcFromSql(resp *[]db.SettingsFindForUpdateRow) *devkitv1.SettingsFindForUpdateResponse {
	grpcRows := make([]*devkitv1.SettingsFindForUpdateRow, len(*resp))
	for index, v := range *resp {
		grpcRow := &devkitv1.SettingsFindForUpdateRow{
			SettingKey:   v.SettingKey,
			SettingValue: v.SettingValue,
			InputType:    v.InputTypeName,
		}

		grpcRows[index] = grpcRow

	}

	return &devkitv1.SettingsFindForUpdateResponse{
		Settings: grpcRows,
	}

}
