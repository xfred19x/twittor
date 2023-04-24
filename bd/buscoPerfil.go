package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/xfred19x/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BuscoPerfil busca un perfil en la BD */
func BuscoPerfil(ID string) (models.Usuario, error) {

	//Agregaremos los parametros de timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	//datos de la BD y la colleccion a usar
	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	//aqui grabaremos todo el perfil del usuario a consultar
	var perfil models.Usuario

	//Aqui obtenemos un objeto del tipo ObjectID
	objId, _ := primitive.ObjectIDFromHex(ID)

	//aqui ponemos la condicion para buscarlo en la BD
	condicion := bson.M{
		"_id": objId,
	}

	//aqui buscamos un solo registro segun la condicion, si todo es ok todo se guardara en Perfil ("usuario")
	err := col.FindOne(ctx, condicion).Decode(&perfil)

	//nunca se debe obtener directamente el passwotd, lo dejaremos en vacio
	perfil.Password = ""
	//validamos si hubo algun error
	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		//se retorna un perfil vacio y el error
		return perfil, err
	}

	//se retorna el perfil cargado y error en nulo
	return perfil, nil
}
