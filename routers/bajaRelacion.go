package routers

import (
	"net/http"

	"github.com/xfred19x/twittor/bd"
	"github.com/xfred19x/twittor/models"
)

/*BajaRelacion realiza el borrado de la relacion entre usuarios */
func BajaRelacion(w http.ResponseWriter, r *http.Request) {

	//obtiene de la URL el parametro "id"
	ID := r.URL.Query().Get("id")

	//crea un objeto model relacion vacio
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	//Llama a la rutina BorroRelacion para eliminar el registro
	status, err := bd.BorroRelacion(t)
	//valido si hay error
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar relación "+err.Error(), http.StatusBadRequest)
		return
	}

	//valida si el status de borrar fue correcto
	if !status {
		http.Error(w, "No se ha logrado borrar la relación "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
