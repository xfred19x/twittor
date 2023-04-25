package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/xfred19x/twittor/bd"
	"github.com/xfred19x/twittor/models"
)

/*GraboTweet permite grabar el tweet en la base de datos */
func GraboTweet(w http.ResponseWriter, r *http.Request) {

	//creamos el objeto con la estructura tweet en vacio
	var mensaje models.Tweet

	//obtenemos del Body el mensaje
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	//Validamos si hubo un error al decodificar el body
	if err != nil {
		http.Error(w, "Mensaje inválido "+err.Error(), 400)
		return
	}

	//creamos la sentencia o condicion de registro
	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	//usamos la funcion para insertar tweet
	_, status, err := bd.InsertoTweet(registro)
	//validamos si no hubo error en la inserccion
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}
	//validamos si el estatus de registro es ok
	if !status {
		http.Error(w, "No se ha logrado insertar el Tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
