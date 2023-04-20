package bd

import "golang.org/x/crypto/bcrypt"

//EncriptarPassword es la rutina que me permite encriptar password
func EncriptarPassword(pass string) (string, error) {
	//el costo es la cantidad basado en un algoritmo de 2 elevado al costo
	//mientras mayor sea el costo, sera la cantidad de pasadas que va sobre el texto para encriptarlo
	//mayor costo, mayor seguridad pero tambien demora más
	//para buenas practicas se usa un costo de 6 para usuarios normales y para admi 8
	costo := 8
	//funcion que retorna bytes
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
