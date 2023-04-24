package routers

import (
	"encoding/json"
	"net/http"

	"github.com/xfred19x/twittor/bd"
)

/*VerPerfil permite extraer los valores del Perfil */
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	//vamos a extraer de la URL el parametro "id"
	ID := r.URL.Query().Get("id")
	//validamos si lo encontro
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	//Buscar el perfil con el ID ingresado
	perfil, err := bd.BuscoPerfil(ID)
	//validar si hubo error
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar buscar el registro "+err.Error(), 400)
		return
	}

	//en caso encontrar el perfil restornar en json
	//w.Header().Set("context-type", "application/json")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
