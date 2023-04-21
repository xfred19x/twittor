package jwt

import (
	"time"
	//"jwt" es el alias que le da al paquete
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/xfred19x/twittor/models"
)

/*GeneroJWT genera el encriptado con JWT */
func GeneroJWT(t models.Usuario) (string, error) {

	//Creando la clave privada con un slice de bytes
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")

	//Vamos a poner la lista del privilegios que se grabará el payload con MapClaims
	//No podemos enviar directamente el modelo "Usuario" ya que tiene el password en su estructura asi que solo enviamos los campos necesarios
	//El campo "exp" es de expiración
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	//NewWithClaims obteneremos un nuevo Json Web Token con el payload
	//manejaremos el algoritmo HS256 para encriptar
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	//SignedString hará que firmemos el token generado con mi clave "MastersdelDesarrollo_grupodeFacebook"
	tokenStr, err := token.SignedString(miClave)
	//Validamos si hubo algun error
	if err != nil {
		//retornara el tokenStr vacio por el error la descripcion del error
		return tokenStr, err
	}

	//retornara el tokenStr en string y con el error nulo
	return tokenStr, nil
}
