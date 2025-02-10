package response

import (
	"server/model/database"

	"github.com/gofrs/uuid"
)

type Login struct {
	User                 database.User `json:"user"`
	AccessToken          string        `json:"access_token"`
	AccessTokenExpiresAt int64         `json:"access_token_expires_at"`
}

type UserCard struct {
	UUID      uuid.UUID `json:"uuid"`
	Username  string    `json:"username"`
	Avatar    string    `json:"avatar"`
	Address   string    `json:"address"`
	Signature string    `json:"signature"`
}

type UserChart struct {
	DateList     []string `json:"date_list"`
	LoginData    []int    `json:"login_data"`
	RegisterData []int    `json:"register_data"`
}
