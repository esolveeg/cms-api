package usecase

import (
	"context"
	devkitv1 "github.com/esolveeg/cms-api/proto_gen/devkit/v1"
)

func (s *PublicUsecase) SettingsUpdate(ctx context.Context, req *devkitv1.SettingsUpdateRequest) error {
	params := s.adapter.SettingsUpdateSqlFromGrpc(req)
	err := s.repo.SettingsUpdate(ctx, params)
	if err != nil {
		return err
	}
	return nil

}

func (u *PublicUsecase) SettingsFindForUpdate(ctx context.Context, req *devkitv1.SettingsFindForUpdateRequest) (*devkitv1.SettingsFindForUpdateResponse, error) {
	settings, err := u.repo.SettingsFindForUpdate(ctx)

	if err != nil {
		return nil, err
	}
	resp := u.adapter.SettingsFindForUpdateGrpcFromSql(settings)

	return resp, nil
}
