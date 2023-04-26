package routers

import (
	"encoding/json"
	"net/http"

	"github.com/xfred19x/twittor/bd"
	"github.com/xfred19x/twittor/models"
)

/*ConsultaRelacion chequea si hay relacion entre 2 usuarios */
func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {

	//obtengo de la URL el parametro "id"
	ID := r.URL.Query().Get("id")

	//creamos un objeto models relacion
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	//creamos un objeto model RespuestaConsultaRelacion quien tendra el estatus si tiene o no una relacion
	var resp models.RespuestaConsultaRelacion

	//consulto la rutina ConsultoRelacion de la BD
	status, err := bd.ConsultoRelacion(t)
	//Valido si hay algun error
	if err != nil || !status {
		//si ocurre algun error o el estatus de la consulta es false, asignar false en el objeto model RespuestaConsultaRelacion
		resp.Status = false
	} else {
		//si la consulta es OK, asignar true en el objeto model RespuestaConsultaRelacion
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
