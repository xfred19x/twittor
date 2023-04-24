package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/xfred19x/twittor/bd"
	"github.com/xfred19x/twittor/models"
)

/*Email valor de Email usado en todos los EndPoints */
var Email string

/*IDUsuario es el ID devuelto del modelo, que se usará en todos los EndPoints */
var IDUsuario string

/*ProcesoToken proceso token para extraer sus valores */
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {

	//Se crea el slide con la palabra clave que se encripto el token
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")

	//se crea un claim para el retorno
	claims := &models.Claim{}

	//se crea la variable y se usa la funcion "Split" para separar la palabra "Bearer" con el token
	splitToken := strings.Split(tk, "Bearer")
	//se valida que el token al menos tenga 2 elementos
	if len(splitToken) != 2 {
		//en caso no tenga elementos requeridos
		//retorna el claims, false por estar mal, el ID con el mensaje vacio y el error con su mensaje personalizado
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	//se convierte el token para quitarle los espacios y asignarle el valor de token real
	tk = strings.TrimSpace(splitToken[1])

	//usaremos la funcion ParseWithClaims, para parsear el token dentro de los claims y de esa forma validar el token
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	//validaremos si hubo algun error
	if err == nil {

		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			//una vez que lo valido obtenemos el email y el ID para que sea usado para todos los demás endpoint al ser variables globales
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		//se retorna el modelo claims, el "true" por encontrar en la BD el email, el ID y el error en nulo
		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		//retorna el claims, false por estar mal, el ID con el mensaje vacio y el error con su mensaje personalizado
		return claims, false, string(""), errors.New("token Inválido")
	}

	//para cerrar todo el circulo
	//Si todo es OK, retorna el claims,el false por no haber error, el ID con el mensaje vacio y el error con su mensaje personalizado
	return claims, false, string(""), err
}
