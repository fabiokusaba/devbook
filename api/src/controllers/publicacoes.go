package controllers

import (
	"encoding/json"
	"github.com/fabiokusaba/devbook/api/src/autenticacao"
	"github.com/fabiokusaba/devbook/api/src/banco"
	"github.com/fabiokusaba/devbook/api/src/modelos"
	"github.com/fabiokusaba/devbook/api/src/repositorios"
	"github.com/fabiokusaba/devbook/api/src/respostas"
	"io/ioutil"
	"net/http"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao modelos.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publicacao.AutorID = usuarioID

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao.ID, erro = repositorio.Criar(publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, publicacao)
}

func BuscarPublicacoes(w http.ResponseWriter, r *http.Request)   {}
func BuscarPublicacao(w http.ResponseWriter, r *http.Request)    {}
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {}
func DeletarPublicacao(w http.ResponseWriter, r *http.Request)   {}
