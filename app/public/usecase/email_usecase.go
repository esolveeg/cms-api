package usecase

import (
	"context"

	devkitv1 "github.com/esolveeg/cms-api/proto_gen/devkit/v1"
)

func (s *PublicUsecase) SendEmail(ctx context.Context, req *devkitv1.SendEmailRequest) (*devkitv1.SendEmailResponse, error) {
	params := s.adapter.SendEmailResendFromGrpc(req)
	resp, err := s.resendClient.SendEmail(&params)
	if err != nil {
		return nil, err
	}
	return &devkitv1.SendEmailResponse{
		Id: resp.Id,
	}, nil

}
