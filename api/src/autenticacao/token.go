package autenticacao

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fabiokusaba/devbook/api/src/config"
)

func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	return token.SignedString([]byte(config.SecretKey))
}
