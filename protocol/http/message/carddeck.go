package message

type UploadCardDeckRequest struct {
	Owner           uint     `json:"owner"`
	RequiredPackage []string `json:"required_package"`
	Cards           []uint   `json:"cards"`
	Characters      []uint   `json:"characters"`
}

type UploadCardDeckResponse struct {
	ID              uint     `json:"id"`
	Owner           uint     `json:"owner"`
	RequiredPackage []string `json:"required_package"`
	Cards           []uint   `json:"cards"`
	Characters      []uint   `json:"characters"`
}

type UpdateCardDeckRequest struct {
	Owner           uint     `json:"owner"`
	RequiredPackage []string `json:"required_package"`
	Cards           []uint   `json:"cards"`
	Characters      []uint   `json:"characters"`
}

type UpdateCardDeckResponse struct {
	ID              uint     `json:"id"`
	Owner           uint     `json:"owner"`
	RequiredPackage []string `json:"required_package"`
	Cards           []uint   `json:"cards"`
	Characters      []uint   `json:"characters"`
}

type QueryCardDeckResponse struct {
	ID              uint     `json:"id"`
	Owner           uint     `json:"owner"`
	RequiredPackage []string `json:"required_package"`
	Cards           []uint   `json:"cards"`
	Characters      []uint   `json:"characters"`
}
