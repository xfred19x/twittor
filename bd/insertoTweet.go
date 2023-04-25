package bd

import (
	"context"
	"time"

	"github.com/xfred19x/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoTweet graba el Tweet en la BD */
func InsertoTweet(t models.GraboTweet) (string, bool, error) {

	//configura el timeout de conexion a la BD
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//configura el nombre de la BD y coleccion
	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	//se crea la condicion o sentencia de registro
	registro := bson.M{
		"userid":  t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}

	//Insertamos el registro
	result, err := col.InsertOne(ctx, registro)
	//validamos si hubo error
	if err != nil {
		//retornamos el ID en vacio, false por el error y la descripcion del error
		return "", false, err
	}

	//aqui obtenemos el ID del resultado
	objID, _ := result.InsertedID.(primitive.ObjectID)
	//retorna el ID, true de todo el proceso OK y el error en nulo.
	return objID.String(), true, nil
}
