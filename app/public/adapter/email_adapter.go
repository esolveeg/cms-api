package adapter

import (
	devkitv1 "github.com/esolveeg/cms-api/proto_gen/devkit/v1"
	"github.com/resend/resend-go/v2"
)

func (a *PublicAdapter) SendEmailResendFromGrpc(req *devkitv1.SendEmailRequest) resend.SendEmailRequest {
	var tags []resend.Tag
	for _, tag := range req.Tags {
		tags = append(tags, resend.Tag{Name: tag.Key, Value: tag.Value})
	}

	var attachments []*resend.Attachment
	for _, attachment := range req.Attachments {
		attachments = append(attachments, &resend.Attachment{
			Filename:    attachment.Filename,
			ContentType: attachment.ContentType,
			Content:     attachment.Content,
		})
	}

	return resend.SendEmailRequest{
		From:        req.From,
		To:          req.To,
		Subject:     req.Subject,
		Bcc:         req.Bcc,
		Cc:          req.Cc,
		ReplyTo:     req.ReplyTo,
		Html:        req.Html,
		Text:        req.Text,
		Tags:        tags,
		Attachments: attachments,
		Headers:     req.Headers,
		ScheduledAt: req.ScheduledAt,
	}
}
