package routers

import (
	"encoding/json"
	"net/http"

	"github.com/xfred19x/twittor/bd"
	"github.com/xfred19x/twittor/models"
)

/*ModificarPerfil modifica el perfil de usuario */
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	//creamos la variable que tendra los valores a modificar
	var t models.Usuario

	//obtenemos del body y lo guardamos en la variable "t" de usuarios
	err := json.NewDecoder(r.Body).Decode(&t)
	//valida si ocurre un error
	if err != nil {
		http.Error(w, "Datos Incorrectos "+err.Error(), 400)
		return
	}

	//creamos la variable que nos indicara si modifico bien o no los datos
	var status bool

	//Modificamos el Registro del Perfil de Usuario
	status, err = bd.ModificoRegistro(t, IDUsuario)
	//validaremos si ocurrio algun error al actualizar
	if err != nil {
		http.Error(w, "Ocurrión un error al intentar modificar el registro. Reintente nuevamente "+err.Error(), 400)
		return
	}

	//Si no actualizo correctamente se envia error
	if !status {
		http.Error(w, "No se ha logrado modificar el registro del usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
