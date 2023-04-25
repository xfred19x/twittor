package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BorroTweet borra un tweet determinado */
func BorroTweet(ID string, UserID string) error {

	//configuramos el timeout de conexion
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	//configuramos el nombre de la BD y la coleccion
	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	//convertimos el ID string en objetID
	objID, _ := primitive.ObjectIDFromHex(ID)

	//creamos la sentencia de condicion para la BD
	condicion := bson.M{
		"_id":    objID,
		"userid": UserID,
	}

	//usamos la funcion eliminar un registro de mongo
	_, err := col.DeleteOne(ctx, condicion)
	//para cualquier error retornarlo
	return err
}
