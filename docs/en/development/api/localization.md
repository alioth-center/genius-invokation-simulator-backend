Here is the GISB localization related services, including the acquisition of language package, translation of specific words and other interfaces

# Variables and arguments

+ `:language_pack_id`/`{language_pack_id}`: The ID of the language package and string type you want to get are registered by the mod provider

# Notes

This interface uses the hard-coded ID of the deterministic language, which is defined as follows and increments from top to bottom, starting at 0:

```go
type Language uint

const (
	ChineseSimplified  Language = iota // ChineseSimplified 简体中文
	ChineseTraditional                 // ChineseTraditional 繁體中文
	English                            // English
	French                             // French Français
	German                             // German Deutsch
	Japanese                           // Japanese 日本語
	Korean                             // Korean 한국어
	Russian                            // Russian Русский язык
	Unknown                            // Unknown ያንሽቤ
)
```

All interfaces in this document do not require authentication

# Getting language packs

## The request method and address

```plain text
GET {BaseURL}/localization/language_pack/:language_pack_id
```

## Request/response

<!-- panels:start -->

<!-- div:left-panel -->

### Request

This interface does not need to provide a request body

#### Sample request

```http
GET /localization/language_pack/:language_pack_id HTTP/1.1
Host: {BaseURL}
```

<!-- div:right-panel -->

### Response

This interface will query and return the language pack based on the language pack ID you give it

#### Response body definition

```go
type LocalizationQueryResponse struct {
    LanguagePack struct{
        Languages          map[uint]map[string]string `json:"languages"`
        SupportedLanguages []enum.Language            `json:"supported_languages"`
    } `json:"language_pack"`
}
```

#### Sample response

```json
{
  "language_pack": {
    "languages": {
      "0": {
        "keqing": "刻晴",
        "zhongli": "钟离"
      },
      "3": {
        "keqing": "keqing",
        "zhongli": "zhongli"
      }
    },
    "supported_languages": [
      3,
      0
    ]
  }
}
```

<!-- panels:end -->

# Translate the given word

## The request method and address

```plain text
GET {BaseURL}/localization/translate?target_language={target_language}&language_package={language_pack_id}
```

## Query parameters

+ `target_language`: 目标语言
+ `language_package`: 词语所在的翻译包ID

## Request/response

<!-- panels:start -->

<!-- div:left-panel -->

### Request

You need to provide the words you want to look up

#### Request body definition

```go
type TranslationRequest struct {
    Words []string `json:"words"`
}
```

#### Sample request

```http
GET /localization/translate?target_language={target_language}&language_package={language_pack_id} HTTP/1.1
Host: {BaseURL}
Content-Type: application/json
Content-Length: 40

{
    "words": ["keqing", "zhongli"]
}
```

<!-- div:right-panel -->

### Response

This interface will try to translate the word you give it with the required language pack into the language you specify

#### Response body definition

```go
type TranslationResponse struct {
    Translation map[string]string `json:"translation"`
}
```

#### Sample response

```json
{
  "translation": {
    "keqing": "刻晴",
    "zhongli": "钟离"
  }
}
```

<!-- panels:end -->