package bd

import (
	"context"
	"time"

	"github.com/xfred19x/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultoRelacion consulta la relacion entre 2 usuarios */
func ConsultoRelacion(t models.Relacion) (bool, error) {

	//configurando timeout de conexion a la BD
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//configurando el nombre de la BD y la colleccion
	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	//crea la condicion de la consulta para la BD
	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	//crea un objeto de moldeo relacion en vacio
	var resultado models.Relacion

	//con la funcion FindOne obtiene el registro de la consulta en resultado
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	//valida si hay error
	if err != nil {
		//retorna "false" si hubo un error y la descripcion
		return false, err
	}

	//retorna "true" si todo es OK y el error nulo
	return true, nil
}
