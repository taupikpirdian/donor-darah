package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func (us *userUsecase) Login(c context.Context, dtoUser *domain.DtoRequestLogin) (t *domain.Auth, err error) {
	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	dataUser, errEntity := domain.NewUserLogin(dtoUser)
	if errEntity != nil {
		return nil, errEntity
	}

	dataUserDb, errUser := us.userRepo.FindUser(ctx, dataUser)
	if errUser != nil {
		return nil, errUser
	}

	// compare the user-entered password with the stored hashed password
	errCompare := bcrypt.CompareHashAndPassword([]byte(dataUserDb.Password), []byte(dtoUser.Password))
	if errCompare != nil {
		return nil, errCompare
	}
	// Set custom claims
	idConv, errConv := strconv.ParseInt(dataUserDb.Id, 10, 64)
	if errConv != nil {
		return nil, errConv
	}

	claims := &domain.JwtCustomClaims{
		Id:   idConv,
		Name: dataUserDb.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey := viper.GetString(`jwt.key`)
	// Generate encoded token and send it as response.
	sign, errSign := token.SignedString([]byte(jwtKey))
	if errSign != nil {
		return nil, errSign
	}

	t, err = domain.SetToken(sign)
	if err != nil {
		return nil, err
	}

	return t, nil
}
