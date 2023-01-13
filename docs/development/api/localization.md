此处为GISB的本地化相关服务，包含语言包的获取、特定词语的翻译等接口

# 变量与参数

+ `:language_pack_id`/`{language_pack_id}`: 您要获取的语言包的ID，字符串类型，由mod提供者注册

# 说明

本接口使用了硬编码确定语言的ID，其定义如下，从上到下值递增，从0开始：

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

本文档的所有接口均不需要鉴权

# 获取语言包

## 请求方法与地址

```plain text
GET {BaseURL}/localization/language_pack/:language_pack_id
```

## 请求/响应

<!-- panels:start -->

<!-- div:left-panel -->

### 请求

本接口不需要提供请求体

#### 请求示例

```http
GET /localization/language_pack/:language_pack_id HTTP/1.1
Host: {BaseURL}
```

<!-- div:right-panel -->

### 响应

此接口将根据您给定的语言包ID，查询并返回语言包

#### 响应体定义

```go
type LocalizationQueryResponse struct {
    LanguagePack struct{
        Languages          map[uint]map[string]string `json:"languages"`
        SupportedLanguages []enum.Language            `json:"supported_languages"`
    } `json:"language_pack"`
}
```

#### 响应示例

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

# 翻译给定词语

## 请求方法与地址

```plain text
GET {BaseURL}/localization/translate?target_language={target_language}&language_package={language_pack_id}
```

## 查询参数

+ `target_language`: 目标语言
+ `language_package`: 词语所在的翻译包ID

## 请求/响应

<!-- panels:start -->

<!-- div:left-panel -->

### 请求

您需要提供需要查询的词语

#### 请求体定义

```go
type TranslationRequest struct {
    Words []string `json:"words"`
}
```

#### 请求示例

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

### 响应

此接口将根据您给定的词语与要求的语言包，尝试将其翻译为您指定的语言

#### 响应体定义

```go
type TranslationResponse struct {
    Translation map[string]string `json:"translation"`
}
```

#### 响应示例

```json
{
  "translation": {
    "keqing": "刻晴",
    "zhongli": "钟离"
  }
}
```

<!-- panels:end -->