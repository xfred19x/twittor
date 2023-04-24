package bd

import (
	"github.com/xfred19x/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin realiza el chequeo de login a la BD */
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	//validamos si ya existe usuario
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)

	//se valido y no encontro el usuario
	if !encontrado {
		//retornara datos del usuario vacio, con el status "false"
		return usu, false
	}

	//creamos un slide de byte para que trabaje con el bcrypt
	passwordBytes := []byte(password)
	//se obtiene el slide con password obtenido de la BD
	passwordBD := []byte(usu.Password)

	//la funcion CompareHashAndPassword, valida el password ingresado con el password encriptado de la BD
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	//validamos si hubo algun error
	if err != nil {
		//retornara datos del usuario vacio, con el status "false"
		return usu, false
	}

	//retornara datos del usuario, con el status "true"
	return usu, true
}
