package routers

import (
	"net/http"

	"github.com/xfred19x/twittor/bd"
)

/*EliminarTweet permite borrar un Tweet determinado */
func EliminarTweet(w http.ResponseWriter, r *http.Request) {

	//obtenemos del URL el parametro "id"
	ID := r.URL.Query().Get("id")
	//validamos que el "id" tenga contenido
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	//usamos la funcion BD para borrar tweet
	err := bd.BorroTweet(ID, IDUsuario)
	//validamos si hubo un error
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
