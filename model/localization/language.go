package localization

import "github.com/sunist-c/genius-invokation-simulator-backend/enum"

type MultipleLanguagePack struct {
	Languages          map[enum.Language]map[string]string `json:"languages"`
	SupportedLanguages []enum.Language                     `json:"supported_languages"`
}

// LanguagePack 语言包接口
type LanguagePack interface {
	Translate(original string, destLanguage enum.Language) (ok bool, result string)
	Pack() MultipleLanguagePack
}

// languagePack 语言包
type languagePack struct {
	translations map[enum.Language]map[string]string
	languages    []enum.Language
}

func (l languagePack) Translate(original string, destLanguage enum.Language) (ok bool, result string) {
	if pack, hasLanguage := l.translations[destLanguage]; !hasLanguage {
		return false, ""
	} else {
		result, ok = pack[original]
		return ok, result
	}
}

func (l languagePack) Pack() MultipleLanguagePack {
	pack := MultipleLanguagePack{
		Languages:          l.translations,
		SupportedLanguages: l.languages,
	}

	return pack
}

func NewLanguagePack(translations map[enum.Language]map[string]string) LanguagePack {
	pack := languagePack{
		translations: translations,
		languages:    []enum.Language{},
	}

	for language := range translations {
		pack.languages = append(pack.languages, language)
	}

	return pack
}
