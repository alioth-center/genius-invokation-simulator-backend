package model

import "github.com/sunist-c/genius-invokation-simulator-backend/model/localization"

type LocalizationQueryResponse struct {
	LanguagePack localization.MultipleLanguagePack `json:"language_pack"`
}

type TranslationRequest struct {
	Words []string `json:"words"`
}

type TranslationResponse struct {
	Translation map[string]string `json:"translation"`
}
