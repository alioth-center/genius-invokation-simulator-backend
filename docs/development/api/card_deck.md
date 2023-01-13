此处为GISB的卡组相关服务，包含卡组的新增、修改、获取、删除等接口

# 变量与参数

+ `:card_deck_id`: 表示卡组的ID
+ `{gisb_token_id_value}`: 表示token_id在cookie中的值，由登录接口获取
+ `{gisb_token_value}`: 表示token在cookie中的键，由登录接口获取

# 说明

本文的所有接口均需要鉴权，您需要在请求中附带您的token

# 上传卡组

## 请求方法与地址

```plain text
POST {BaseURL}/card_deck
```

## 请求/响应

<!-- panels:start -->

<!-- div:left-panel -->

### 请求

您需要提供所要上传卡组的信息

#### 请求体定义

```go
type UploadCardDeckRequest struct {
    Owner           uint     `json:"owner"`
    RequiredPackage []string `json:"required_package"`
    Cards           []uint   `json:"cards"`
    Characters      []uint   `json:"characters"`
}
```

#### 请求示例

```http
POST /card_deck HTTP/1.1
Host: {BaseURL}
Content-Type: application/json
Cookie: gisb_token={gisb_token_value}; gisb_token_id={gisb_token_id_value}
Content-Length: 98

{
    "owner": 1,
    "required_package": ["base"],
    "cards": [1],
    "characters": [1]
}
```

<!-- div:right-panel -->

### 响应

此接口将尝试把您提供的卡组上传至持久化模块，并同步更新玩家信息

#### 响应体定义

```go
type UploadCardDeckResponse struct {
    ID              uint     `json:"id"`
    Owner           uint     `json:"owner"`
    RequiredPackage []string `json:"required_package"`
    Cards           []uint   `json:"cards"`
    Characters      []uint   `json:"characters"`
}
```

#### 响应示例

```json
{
  "id": 3,
  "owner": 1,
  "required_package": [
    "base"
  ],
  "cards": [
    1
  ],
  "characters": [
    1
  ]
}
```

<!-- panels:end -->

# 获取卡组

## 请求方法与地址

```plain text
GET {BaseURL}/card_deck/:card_deck_id
```

## 请求/响应

<!-- panels:start -->

<!-- div:left-panel -->

### 请求

本接口不需要提供请求体

#### 请求示例

```http
GET /card_deck/:card_deck_id HTTP/1.1
Host: {BaseURL}
Cookie: gisb_token={gisb_token_value}; gisb_token_id={gisb_token_id_value}
```

<!-- div:right-panel -->

### 响应

此接口将查询并返回您所给定id的卡组

#### 响应体定义

```go
type QueryCardDeckResponse struct {
    ID              uint     `json:"id"`
    Owner           uint     `json:"owner"`
    RequiredPackage []string `json:"required_package"`
    Cards           []uint   `json:"cards"`
    Characters      []uint   `json:"characters"`
}
```

#### 响应示例

```json
{
  "id": 3,
  "owner": 1,
  "required_package": [
    "base"
  ],
  "cards": [
    1
  ],
  "characters": [
    1
  ]
}
```

<!-- panels:end -->

# 更新卡组

## 请求方法与地址

```plain text
PUT {BaseURL}/card_deck/:card_deck_id
```

## 请求/响应

<!-- panels:start -->

<!-- div:left-panel -->

### 请求

您需要提供所要更新卡组的信息

#### 请求体定义

```go
type UpdateCardDeckRequest struct {
    Owner           uint     `json:"owner"`
    RequiredPackage []string `json:"required_package"`
    Cards           []uint   `json:"cards"`
    Characters      []uint   `json:"characters"`
}
```

#### 请求示例

```http
PUT /card_deck/:card_deck_id HTTP/1.1
Host: {BaseURL}
Content-Type: application/json
Cookie: gisb_token={gisb_token_value}; gisb_token_id={gisb_token_id_value}
Content-Length: 101

{
    "owner": 1,
    "required_package": ["enhance"],
    "cards": [1],
    "characters": [1]
}
```

<!-- div:right-panel -->

### 响应

此接口将根据您所给定的id尝试更新卡组，并将结果返回

#### 响应体定义

```go
type UpdateCardDeckResponse struct {
    ID              uint     `json:"id"`
    Owner           uint     `json:"owner"`
    RequiredPackage []string `json:"required_package"`
    Cards           []uint   `json:"cards"`
    Characters      []uint   `json:"characters"`
}
```

#### 响应示例

```json
{
  "id": 3,
  "owner": 1,
  "required_package": [
    "enhance"
  ],
  "cards": [
    1
  ],
  "characters": [
    1
  ]
}
```

<!-- panels:end -->

# 删除卡组

## 请求方法与地址

```plain text
DELETE {BaseURL}/card_deck/:card_deck_id
```

## 请求/响应

<!-- panels:start -->

<!-- div:left-panel -->

### 请求

本接口不需要提供请求体

#### 请求示例

```http
DELETE /card_deck/:card_deck_id HTTP/1.1
Host: {BaseURL}
Cookie: gisb_token={gisb_token_value}; gisb_token_id={gisb_token_id_value}
```

<!-- div:right-panel -->

### 响应

此接口将根据您给定的id尝试删除卡组

#### 响应体定义

本接口没有响应体，返回HTTP Status 200即表明删除成功

<!-- panels:end -->