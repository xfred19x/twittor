package bd

import (
	"context"
	"time"

	"github.com/xfred19x/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ModificoRegistro permite modificar el perfil del usuario */
func ModificoRegistro(u models.Usuario, ID string) (bool, error) {

	//configuramos el tiempo de timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//definimos el nombre de la bd y la colleccion
	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	//la variable registro vacío con make para crear un mapa del tipo string con valores interface
	//esto es para tener distintas alternativas para armar el registro de actualizacion
	registro := make(map[string]interface{})

	//validamos que los campos tengan valores
	//al final grabo el clave valor en el "registro"
	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}
	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos
	}
	registro["fechaNacimiento"] = u.FechaNacimiento
	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}
	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}
	if len(u.Ubicacion) > 0 {
		registro["ubicacion"] = u.Ubicacion
	}
	if len(u.SitioWeb) > 0 {
		registro["sitioWeb"] = u.SitioWeb
	}

	//creamos la condicion de registro y actualizacion con el variable "registro"
	updtString := bson.M{
		"$set": registro,
	}

	//convierte mi ID en string a ObjectID
	objID, _ := primitive.ObjectIDFromHex(ID)
	//actualizamos la base de mongo con el filtro que sera el "objID"
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	//usamos la funcion UpdateOne para actualizar un registro
	_, err := col.UpdateOne(ctx, filtro, updtString)
	//Valida si hay error
	if err != nil {
		//retorna "false" por algun problema mas el error
		return false, err
	}

	//retorna el "true" porque todo estuvo bien y en el error nulo
	return true, nil
}
