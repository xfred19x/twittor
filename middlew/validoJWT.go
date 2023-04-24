package middlew

import (
	"net/http"

	"github.com/xfred19x/twittor/routers"
)

/*ValidoJWT premite validar el JWT que nos viene en la peticion */
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//llamaremos la rutina ProcesoToken que verificara si el token es valido o no
		//el campo a validar del header es "Authorization"
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		//validaremos si hubo algun error en la validacion
		if err != nil {
			//se enviara como respuesta el codstatus de la respuesta
			http.Error(w, "Error en el TOken ! "+err.Error(), http.StatusBadRequest)
			return
		}

		//pasamos los objetos para la siguiente cadena
		next.ServeHTTP(w, r)
	}
}
