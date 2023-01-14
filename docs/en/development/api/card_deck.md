This is the card deck related service of GISB, including interfaces for adding, modifying, acquiring, and deleting card decks.

# Variables and arguments

+ `:card_deck_id`: Represents the ID of the card deck
+ `{gisb_token_id_value}`: Represents the value of the token_id in the cookie, obtained by the login interface
+ `{gisb_token_value}`: Representing the key of token in the cookie, obtained by the login interface

# Explain

All interfaces in this topic require authentication, and you need to include your token in the request

# Upload the card deck

## Request method and address

```plain text
POST {BaseURL}/card_deck
```

## Request/response

<!-- panels:start -->

<!-- div:left-panel -->

### Request

You'll need to provide information about the card deck you want to upload

#### Request body definition

```go
type UploadCardDeckRequest struct {
    Owner           uint     `json:"owner"`
    RequiredPackage []string `json:"required_package"`
    Cards           []uint   `json:"cards"`
    Characters      []uint   `json:"characters"`
}
```

#### Request example

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

### Response

This interface will attempt to upload the card deck you provided to the persistence module and update the player information synchronously


#### Response body definition

```go
type UploadCardDeckResponse struct {
    ID              uint     `json:"id"`
    Owner           uint     `json:"owner"`
    RequiredPackage []string `json:"required_package"`
    Cards           []uint   `json:"cards"`
    Characters      []uint   `json:"characters"`
}
```

#### Sample response 

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

# Acquisition card beck

## The request method and address

```plain text
GET {BaseURL}/card_deck/:card_deck_id
```

## Request/response

<!-- panels:start -->

<!-- div:left-panel -->

### Request

This interface does not need to provide a request body

#### Sample request

```http
GET /card_deck/:card_deck_id HTTP/1.1
Host: {BaseURL}
Cookie: gisb_token={gisb_token_value}; gisb_token_id={gisb_token_id_value}
```

<!-- div:right-panel -->

### Response

This interface will query and return the card deck with the given id

#### Response body definition

```go
type QueryCardDeckResponse struct {
    ID              uint     `json:"id"`
    Owner           uint     `json:"owner"`
    RequiredPackage []string `json:"required_package"`
    Cards           []uint   `json:"cards"`
    Characters      []uint   `json:"characters"`
}
```

#### Sample response

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

# Update card deck

## The request method and address

```plain text
PUT {BaseURL}/card_deck/:card_deck_id
```

## Request/response

<!-- panels:start -->

<!-- div:left-panel -->

### Request

You need to provide information about the card deck you want to update

#### Request body definition

```go
type UpdateCardDeckRequest struct {
    Owner           uint     `json:"owner"`
    RequiredPackage []string `json:"required_package"`
    Cards           []uint   `json:"cards"`
    Characters      []uint   `json:"characters"`
}
```

#### Sample request

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

### Response

This interface will attempt to update the card deck based on the id you have given and return the result

#### Response body definition

```go
type UpdateCardDeckResponse struct {
    ID              uint     `json:"id"`
    Owner           uint     `json:"owner"`
    RequiredPackage []string `json:"required_package"`
    Cards           []uint   `json:"cards"`
    Characters      []uint   `json:"characters"`
}
```

#### Sample request

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

# Deleting card deck

## The request method and address

```plain text
DELETE {BaseURL}/card_deck/:card_deck_id
```

## Request/response

<!-- panels:start -->

<!-- div:left-panel -->

### Request

This interface does not need to provide a request body

#### Sample request

```http
DELETE /card_deck/:card_deck_id HTTP/1.1
Host: {BaseURL}
Cookie: gisb_token={gisb_token_value}; gisb_token_id={gisb_token_id_value}
```

<!-- div:right-panel -->

### Response

This interface will attempt to remove the card deck based on the id you give it

#### Response body definition

This interface does not have a response body, so returning HTTP Status 200 indicates that the deletion was successful

<!-- panels:end -->