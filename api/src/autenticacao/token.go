package autenticacao

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)

	// Validando o token recebido com a chave de verificação através do parse
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return erro
	}

	// Verificando a existência dos claims
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token inválido")
}

func ExtrairUsuarioId(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)

	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioId, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return usuarioId, nil
	}

	return 0, errors.New("Token inválido")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	// Verificando o formato do token -> "Bearer <token>"
	if len(strings.Split(token, " ")) == 2 {
		// Retornando apenas o token, no caso a segunda palavra -> <token>
		return strings.Split(token, " ")[1]
	}

	// Retornando uma string em branco no caso do token não estar no formato válido
	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
