package bd

import (
	"context"
	"time"

	"github.com/xfred19x/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ChequeoYaExisteUsuario recibe un email de parametro y chequea si ya esta en la BD
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	//Haremos una condicion en MongoBD
	//El bson.M es un formato que admite y retorna un mapstring json
	condicion := bson.M{"email": email}

	var resultado models.Usuario

	//la funcion FindOne busca solo un registro
	//luego que lo encuentre usamos el Decode para que lo convierta a JSON
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	//se validara si hubo algun error
	if err != nil {
		//se retorna el resultado, el false de no Existoso y el ID vació
		return resultado, false, ID
	}

	//se retorna el resultado, el  true de OK y el ID
	return resultado, false, ID
}
