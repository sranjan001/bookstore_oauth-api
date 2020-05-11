package access_token

import (
	"github.com/sranjan001/bookstore_oauth-api/src/utils/errors"
	"math/rand"
	"strings"
	"time"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`

	//Used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	//Used for client_credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`

	Scope string `json:"scope"`
}

func (at *AccessTokenRequest) validate() *errors.RestError {
	if at.GrantType != grantTypePassword && at.GrantType != grantTypeClientCredentials {
		return errors.NewBadRequestError("Invalid grant_type parameters")
	}

	//TODO validate for each grant type

	return nil
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	at.AccessToken = b.String() // E.g. "ExcbsVQs"
}

//Web frontend - Client-Id: 123
//Android APP - Client-Id: 234
