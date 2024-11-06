package adapter

import (
	"github.com/esolveeg/cms-api/db"
	devkitv1 "github.com/esolveeg/cms-api/proto_gen/devkit/v1"
)

func (a *PublicAdapter) TranslationsCreateUpdateBulkSqlFromGrpc(req *devkitv1.TranslationsCreateUpdateBulkRequest) *db.TranslationsCreateUpdateBulkParams {
	keys := make([]string, len(req.Records))
	enValues := make([]string, len(req.Records))
	arValues := make([]string, len(req.Records))
	for index, v := range req.Records {
		keys[index] = v.TranslationKey
		enValues[index] = v.EnglishValue
		arValues[index] = v.ArabicValue
	}
	return &db.TranslationsCreateUpdateBulkParams{
		Keys:          keys,
		ArabicValues:  arValues,
		EnglishValues: enValues,
	}
}

func (a *PublicAdapter) TranslationCreateUpdateBulkRowGrpcFromSql(resp *db.TranslationsCreateUpdateBulkRow) *devkitv1.Translation {
	return &devkitv1.Translation{
		TranslationKey: resp.TranslationKey,
		EnglishValue:   resp.EnglishValue,
		ArabicValue:    resp.ArabicValue,
	}
}

func (a *PublicAdapter) TranslationGrpcFromSql(resp *db.Translation) *devkitv1.Translation {
	return &devkitv1.Translation{
		TranslationKey: resp.TranslationKey,
		EnglishValue:   resp.EnglishValue,
		ArabicValue:    resp.ArabicValue,
	}
}

func (a *PublicAdapter) TranslationsCreateUpdateBulkGrpcFromSql(resp []db.TranslationsCreateUpdateBulkRow) devkitv1.TranslationsListResponse {
	translations := make([]*devkitv1.Translation, len(resp))
	for index, t := range resp {
		translations[index] = a.TranslationCreateUpdateBulkRowGrpcFromSql(&t)
	}
	return devkitv1.TranslationsListResponse{
		Translations: translations,
	}
}
func (a *PublicAdapter) TranslationsListGrpcFromSql(resp []db.Translation) devkitv1.TranslationsListResponse {
	translations := make([]*devkitv1.Translation, len(resp))
	for index, t := range resp {
		translations[index] = a.TranslationGrpcFromSql(&t)
	}
	return devkitv1.TranslationsListResponse{
		Translations: translations,
	}
}
