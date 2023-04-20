package bd

import (
	"context"
	"time"

	"github.com/xfred19x/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertoRegistro es la parada final con la BD para insertar los datos del usuario
func InsertoRegistro(u models.Usuario) (string, bool, error) {

	//aqui configuro con el conxteto el tiempo de ejecucion en timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	//el defer se ejecutara al ultimo, si sucede el timeout este realizara un cancel al WithTimeout y el contexto
	defer cancel()

	//Se creara la conexion de bd
	db := MongoCN.Database("twittor")
	//aqui se indica con que coleccion se va trabajar en este caso usuarios
	col := db.Collection("usuarios")

	//se encripta password antes de registrar a la BD
	u.Password, _ = EncriptarPassword(u.Password)

	//Aqui indicaremos que a la coleccion haremos una solo registro con InsertOne
	result, err := col.InsertOne(ctx, u)
	//validaremos si hubo un error al insertar
	if err != nil {
		//retornara vacio, un false de no satisfactorio y el error
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	//retornara el ID del registrod e usuario, el true de OK y el nulo por que no hay error
	return ObjID.String(), true, nil

}
