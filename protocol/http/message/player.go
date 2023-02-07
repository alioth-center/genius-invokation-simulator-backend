package message

import "github.com/sunist-c/genius-invokation-simulator-backend/persistence"

type LoginResponse struct {
	PlayerUID       uint                   `json:"player_uid"`
	Success         bool                   `json:"success"`
	PlayerNickName  string                 `json:"player_nick_name"`
	PlayerCardDecks []persistence.CardDeck `json:"player_card_decks"`
}

type LoginRequest struct {
	PlayerUID uint   `json:"player_uid"`
	Password  string `json:"password"`
}

type RegisterRequest struct {
	NickName string `json:"nick_name"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	PlayerUID      uint   `json:"player_uid"`
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
