package middlew

import (
	"net/http"

	"github.com/xfred19x/twittor/bd"
)

// ChequeoBD es el midlew que me permite conocer el estado de la BD
// los Middleware deben recibir y retornar el mismo tipo de dato
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {

	//debemos retornar una funcion anónima, va tener un proceso adentro
	//validara que aun exista la conexion a la BD
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la Base de Datos", 500)
			return
		}

		//aqui le paso todos los parametros (w, r) al proximo eslabon de la cadena
		next.ServeHTTP(w, r)
	}
}
