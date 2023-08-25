package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func (us *userUsecase) RefreshToken(c context.Context, userId int64) (t *domain.Token, err error) {
	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	dataUser := domain.NewUser4(userId)
	dataUserDb, errUser := us.userRepo.FindUserById(ctx, dataUser)
	if errUser != nil {
		return nil, errUser
	}

	// Set custom claims
	idConv, errConv := strconv.ParseInt(dataUserDb.Id, 10, 64)
	if errConv != nil {
		return nil, errConv
	}

	claims := &domain.JwtCustomClaims{
		Id:   idConv,
		Name: dataUserDb.Name,
		Role: dataUserDb.Role,
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

	return &domain.Token{
		Token: sign,
	}, nil
}
