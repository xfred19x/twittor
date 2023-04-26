package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/xfred19x/twittor/bd"
)

/*ListaUsuarios leo la lista de los usuarios */
func ListaUsuarios(w http.ResponseWriter, r *http.Request) {

	///obtengo de la URL los siguientes parametros
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	//convertimos el valor de "page" a numerico
	pagTemp, err := strconv.Atoi(page)
	//validar error
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	//lo convertimos en numerico int64
	pag := int64(pagTemp)

	//llamo a la rutina LeoUsuariosTodos
	result, status := bd.LeoUsuariosTodos(IDUsuario, pag, search, typeUser)
	//si estatus  es false
	if !status {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
