package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/xfred19x/twittor/bd"
)

/*LeoTweetsSeguidores lee los tweets de todos nuestros seguidores */
func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request) {

	//validamos del URl el parametro "pagina" que tenga valor
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}

	//convertimos a numero el parametro "pagina"
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	//validar si hubo error
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	//consultara la rutina LeoTweetsSeguidores
	respuesta, correcto := bd.LeoTweetsSeguidores(IDUsuario, pagina)
	//valida si es la consulta es correcto
	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
