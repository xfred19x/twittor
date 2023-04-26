package routers

import (
	"net/http"

	"github.com/xfred19x/twittor/bd"
	"github.com/xfred19x/twittor/models"
)

/*AltaRelacion realiza el registro de la relacion entre usuarios */
func AltaRelacion(w http.ResponseWriter, r *http.Request) {

	//vamos a obtener el parametro "id" del URL
	ID := r.URL.Query().Get("id")
	//valida que tenga contenido el parametro "id"
	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	//creamos el objeto con el modelo de relacion
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	//llamamos a la rutina insertaRelacion para registrar la relacion a la BD
	status, err := bd.InsertoRelacion(t)
	//validamos si hubo error
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar relación "+err.Error(), http.StatusBadRequest)
		return
	}

	//validamos que el estatus de la relacion sea correcta
	if !status {
		http.Error(w, "No se ha logrado insertar la relación "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
