package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/xfred19x/twittor/bd"
	"github.com/xfred19x/twittor/jwt"
	"github.com/xfred19x/twittor/models"
)

/*Login realiza el login*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	//Creamos un variable tipo Usuario
	var t models.Usuario

	//vamos a decodificar el body y cargar los datos el la variable usuario "t"
	err := json.NewDecoder(r.Body).Decode(&t)

	//Validamos si hubo un error
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña inválidos "+err.Error(), 400)
		return
	}

	//Validamos si se esta enviando el Email
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido ", 400)
		return
	}

	//Validamos si existe el login y pass
	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if !existe {
		http.Error(w, "Usuario y/o Contraseña Inválidos ", 400)
		return
	}

	//GeneroJWT recibe el documento "usuario" para obtener el token
	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrió un erro al intentar generar el Token correspondiente"+err.Error(), 400)
		return
	}

	//armamos el JSON donde retornamos el token al navegador
	resp := models.RespuesatLogin{
		Token: jwtKey,
	}

	//Estamos seteando en la cabecera el tipo de respuesta JSON
	w.Header().Set("Content-Type", "application/json")
	//Retornamos el status code de creacion
	w.WriteHeader(http.StatusCreated)
	//Retornamos un status de respuesta
	json.NewEncoder(w).Encode(resp)

	//BONUS
	//creacion una variable donde definimos la fecha de expiracion de token
	expirationTime := time.Now().Add(24 * time.Hour)
	//setearemos una cookie en formato JSON
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
