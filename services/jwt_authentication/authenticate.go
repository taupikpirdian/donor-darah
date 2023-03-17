package jwt_authentication

import (
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/golang-jwt/jwt/v4"
)

var APPLICATION_NAME = "DonorDarahSalt"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("qSYWzKQ8-eEqEjWgE-MmAFpg6r")

type AuthJWT struct {
	repoUser   domain.ArticleRepository
	repoAuthor domain.AuthorRepository
}

type AuthJWTInterface interface {
	AuthenticateUser(Phone string, PlainPassword string) (token string, err error)
	AuthenticateAuthor(Phone string, PlainPassword string) (token string, err error)
}

func NewAuthJwt(RepoUser domain.ArticleRepository, RepoAuthor domain.AuthorRepository) AuthJWTInterface {
	return &AuthJWT{
		repoUser:   RepoUser,
		repoAuthor: RepoAuthor,
	}
}

func (aj *AuthJWT) AuthenticateUser(Phone string, PlainPassword string) (token string, err error) {
	return "", nil
}

func (aj *AuthJWT) AuthenticateAuthor(Phone string, PlainPassword string) (token string, err error) {
	return "", nil
}
