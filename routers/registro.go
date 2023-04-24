package routers

import (
	"encoding/json" //libreria que se usa para manejar estructuras en json
	"net/http"

	"github.com/xfred19x/twittor/bd"
	"github.com/xfred19x/twittor/models"
)

// Registro es la funcion para crear en la BD el registro de usuario
func Registro(w http.ResponseWriter, r *http.Request) {

	//creo un modelo de usuario
	var t models.Usuario
	//luego con la funcion NewDecoder se decodifica el Body en el modelo usuario
	//tomar en cuenta que el Body de un Http Request es un stream de un solo uso ya que se destruye en memoria
	err := json.NewDecoder(r.Body).Decode(&t)

	//validamos si se origino algun error por algun body vacio, etc.
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	//Validaremos que se envio correctamente el email
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	//Validaremos que el password es menor a 6
	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
		return
	}

	//Se valida si el email que se esta intentado registrar ya existe
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	//Aqui se realizara el registro
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

	//se validara el status porque a veces mongo retorna vacio
	if !status {
		http.Error(w, "No se ha logrado insertar el registro de usuario ", 400)
		return
	}

	//vamos a deolver por header un http constante que ya estan creadas
	w.WriteHeader(http.StatusCreated)

}
