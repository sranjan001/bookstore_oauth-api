package access_token

import (
	"fmt"
	"github.com/sranjan001/bookstore_oauth-api/src/respository/rest"
	"github.com/sranjan001/bookstore_oauth-api/src/utils/errors"
	"strings"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestError)
	Create(AccessToken) *errors.RestError
	UpdateExpirationTime(AccessToken) *errors.RestError
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestError)
	Create(AccessTokenRequest) (*AccessToken, *errors.RestError)
	UpdateExpirationTime(AccessToken) *errors.RestError
}

type service struct {
	dbRepo   Repository
	userRepo rest.RestUserRepository
}

func NewService(repo Repository, userRepo rest.RestUserRepository) Service {
	return &service{
		dbRepo:   repo,
		userRepo: userRepo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestError) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.dbRepo.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(atr AccessTokenRequest) (*AccessToken, *errors.RestError) {
	if err := atr.validate(); err != nil {
		return nil, err
	}

	//TODO: support both client_credentials and password grant types
	//Authenticate user against the Users Api
	user, err := s.userRepo.LoginUser(atr.Username, atr.Password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("User : ", user)
	//Generate a new access token
	at := GetNewAccessToken(user.Id)
	at.Generate()

	fmt.Println("User accesstoken : ", at)
	//Save the new access token in Cassandra
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestError {
	//if err := at.validate(); err != nil {
	//	return err
	//}
	//return s.dbRepo.UpdateExpirationTime(at)
	return nil
}
