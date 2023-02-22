package message

type UploadCardDeckRequest struct {
	Owner           uint64   `json:"owner"`
	RequiredPackage []string `json:"required_package"`
	Cards           []uint64 `json:"cards"`
	Characters      []uint64 `json:"characters"`
}

type UploadCardDeckResponse struct {
	ID              uint64   `json:"id"`
	Owner           uint64   `json:"owner"`
	RequiredPackage []string `json:"required_package"`
	Cards           []uint64 `json:"cards"`
	Characters      []uint64 `json:"characters"`
}

type UpdateCardDeckRequest struct {
	Owner           uint64   `json:"owner"`
	RequiredPackage []string `json:"required_package"`
	Cards           []uint64 `json:"cards"`
	Characters      []uint64 `json:"characters"`
}

type UpdateCardDeckResponse struct {
	ID              uint64   `json:"id"`
	Owner           uint64   `json:"owner"`
	RequiredPackage []string `json:"required_package"`
	Cards           []uint64 `json:"cards"`
	Characters      []uint64 `json:"characters"`
}

type QueryCardDeckResponse struct {
	ID              uint64   `json:"id"`
	Owner           uint64   `json:"owner"`
	RequiredPackage []string `json:"required_package"`
	Cards           []uint64 `json:"cards"`
	Characters      []uint64 `json:"characters"`
}
