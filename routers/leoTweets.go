package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/xfred19x/twittor/bd"
)

/*LeoTweets Leo los tweets */
func LeoTweets(w http.ResponseWriter, r *http.Request) {

	//obtenemos de la URL el parametro "id"
	ID := r.URL.Query().Get("id")
	//validamos que el "id" tenga contenido
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	//validamos que se enviue el valor de paginar
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}

	//Vamos a convertir pagina string a un valor entero
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	//validamos si hubo error
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	//convertimos el valor numerico a "int64"
	pag := int64(pagina)
	//usamos la funcion LeoTweets, para obtener los tweets con filtros
	respuesta, correcto := bd.LeoTweets(ID, pag)

	//validamos si es correcto
	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
