package message

type CardDeck struct {
	ID               uint64   `json:"id"`
	OwnerUID         uint64   `json:"owner_uid"`
	RequiredPackages []string `json:"required_packages"`
	Cards            []uint64 `json:"cards"`
	Characters       []uint64 `json:"characters"`
}

type LoginResponse struct {
	PlayerUID       uint64     `json:"player_uid"`
	Success         bool       `json:"success"`
	PlayerNickName  string     `json:"player_nick_name"`
	PlayerCardDecks []CardDeck `json:"player_card_decks"`
}

type LoginRequest struct {
	PlayerUID uint64 `json:"player_uid"`
	Password  string `json:"password"`
}

type RegisterRequest struct {
	NickName string `json:"nick_name"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	PlayerUID      uint64 `json:"player_uid"`
	PlayerNickName string `json:"player_nick_name"`
}

type UpdatePasswordRequest struct {
	OriginalPassword string `json:"original_password"`
	NewPassword      string `json:"new_password"`
}

type UpdateNickNameRequest struct {
	Password    string `json:"password"`
	NewNickName string `json:"new_nick_name"`
}

type DestroyPlayerRequest struct {
	Password string `json:"password"`
	Confirm  bool   `json:"confirm"`
}

type DestroyPlayerResponse struct {
	Success bool `json:"success"`
}
