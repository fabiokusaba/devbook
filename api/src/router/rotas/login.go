package rotas

import (
	"net/http"

	"github.com/fabiokusaba/devbook/api/src/controllers"
)

var rotaLogin = Rota{
	URI:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}
