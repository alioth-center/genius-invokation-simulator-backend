此处为GISB的玩家相关服务，包含玩家的注册、登录、更新、注销等接口

# 变量与参数

+ `:player_id`: 表示玩家的UID
+ `{gisb_token_id}`: 表示token_id在cookie中的键，由服务器设置
+ `{gisb_token}`: 表示token在cookie中的键，由服务器进行设置

# 玩家注册

## 请求方法与地址

```plain text
POST {BaseURL}/player
```

## 请求/响应

<!-- panels:start -->

<!-- div:left-panel -->

### 请求

您需要提供注册玩家的注册昵称与密码，您可以将密码随意编码，后端会对密码进行单向加密处理，您只需保证您的密码编码方式不变即可

#### 请求体定义

```go
type RegisterRequest struct {
    NickName string `json:"nick_name"`
    Password string `json:"password"`
}
```

#### 请求示例

```http
POST /player HTTP/1.1
Host: {BaseURL}
Content-Type: application/json
Content-Length: 56

{
    "nick_name": "venti",
    "password": "eihei"
}
```

<!-- div:right-panel -->

### 响应

此接口将会把发起者提供的玩家信息尝试写入GISB的持久化模块中，若写入成功，将返回注册玩家的ID与昵称

#### 响应体定义

```go
type RegisterResponse struct {
    PlayerUID      uint   `json:"player_uid"`
    PlayerNickName string `json:"player_nick_name"`
}
```

#### 响应示例

```json
{
    "player_uid": 114514,
    "player_nick_name": "venti"
}
```

<!-- panels:end -->

# 玩家登录

## 请求方法与地址

```plain text
GET {BaseURL}/player/login/:player_id
```

## 路径参数说明

+ `:player_id`: 请求登录的玩家的UID，必须为整数

## 请求/响应

<!-- panels:start -->

<!-- div:left-panel -->

### 请求

您需要提供等物玩家的密码，编码/加密方式需和注册时一致

#### 请求体定义

```go
type LoginRequest struct {
    Password string `json:"password"`
}
```

#### 请求示例

```http
GET /player/login/:player_id HTTP/1.1
Host: {BaseURL}
Content-Type: application/json
Content-Length: 29

{
    "password": "eihei"
}
```

<!-- div:right-panel -->

### 响应

此接口将会把发起者提供的玩家密码尝试登录，若成功会返回玩家的基本信息和保存的卡组信息

#### 响应体定义

```go
type LoginResponse struct {
    PlayerUID       uint                   `json:"player_uid"`
    Success         bool                   `json:"success"`
    PlayerNickName  string                 `json:"player_nick_name"`
    PlayerCardDecks []struct{
        ID               uint     `json:"id"`
        OwnerUID         uint     `json:"owner_uid"`
        RequiredPackages []string `json:"required_packages"`
        Cards            []uint   `json:"cards"`
        Characters       []uint   `json:"characters"`
    } `json:"player_card_decks"`
}
```

#### 响应示例

```json
{
  "player_uid": 114514,
  "success": true,
  "player_nick_name": "venti",
  "player_card_decks": [
    {
      "id": 1,
      "owner_uid": 1,
      "required_packages": ["base"],
      "cards": [1, 1, 2, 3, 5, 8, 13, 21],
      "characters": [2, 3, 5, 7]
    },
    {
      "id": 2,
      "owner_uid": 1,
      "required_packages": ["base"],
      "cards": [11, 45, 14, 19, 19, 810],
      "characters": [5, 6, 7]
    }
  ]
}
```

<!-- panels:end -->

## SetCookie

若登录成功，您将会得到下面两个SetCookie信息：

+ `{gisb_token_id}`: 此次会话的token_id，在访问需要鉴权的接口时需要将其正确设置
+ `{gisb_token}`： 此次会话的token，在访问需要鉴权的接口时需要将其正确设置

# 更新玩家密码

## 请求方法与地址

```plain text
PATCH {BaseURL}/player/:player_id/password
```

## 请求/响应

<!-- panels:start -->

<!-- div:left-panel -->

### 请求

您需要提供需要更改密码的玩家的当前密码和新密码，编码/加密方式需和注册时相同

#### 请求体定义

```go
type UpdatePasswordRequest struct {
    OriginalPassword string `json:"original_password"`
    NewPassword      string `json:"new_password"`
}
```

#### 请求示例

```http
PATCH /player/:player_id/password HTTP/1.1
Host: {BaseURL}
Content-Type: application/json
Content-Length: 68

{
    "original_password": "eihei",
    "new_password": "hello"
}
```

<!-- div:right-panel -->

### 响应

此接口没有响应体，HTTP Status 200即表明更新成功

<!-- panels:end -->

# 更新玩家昵称

## 请求方法与地址

```plain text
PATCH {BaseURL}/player/:player_id/nickname
```

## 请求/响应

<!-- panels:start -->

<!-- div:left-panel -->

### 请求

您需要提供需要更新昵称的玩家的当前密码和新昵称，编码/加密方式需和注册时相同

#### 请求体定义

```go
type UpdateNickNameRequest struct {
    Password    string `json:"password"`
    NewNickName string `json:"new_nick_name"`
}
```

#### 请求示例

```http
PATCH /player/:player_id/nickname HTTP/1.1
Host: {BaseURL}
Content-Type: application/json
Content-Length: 65

{
    "password": "eihei",
    "new_nick_name": "maichangde"
}
```

<!-- div:right-panel -->

### 响应

此接口没有响应体，HTTP Status 200即表明更新成功

<!-- panels:end -->

# 注销玩家账号

## 请求方法与地址

```plain text
DELETE {BaseURL}/player/:player_id
```

## 请求/响应

<!-- panels:start -->

<!-- div:left-panel -->

### 请求

您需要提供需要注销账号玩家的密码，编码/加密方式需和注册时相同

#### 请求体定义

```go
type DestroyPlayerRequest struct {
    Password string `json:"password"`
    Confirm  bool   `json:"confirm"`
}
```

#### 请求示例

```http
DELETE /player/:player_id HTTP/1.1
Host: {BaseURL}
Content-Type: application/json
Content-Length: 51

{
    "password": "eihei",
    "confirm": true
}
```

<!-- div:right-panel -->

### 响应

此接口会返回注销是否成功

#### 响应体定义

```go
type DestroyPlayerResponse struct {
    Success bool `json:"success"`
}
```

#### 响应示例

```json
{
  "success": true
}
```

<!-- panels:end -->