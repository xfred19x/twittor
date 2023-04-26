package bd

import (
	"context"
	"time"

	"github.com/xfred19x/twittor/models"
)

/*BorroRelacion borra la relacion en la BD */
func BorroRelacion(t models.Relacion) (bool, error) {

	//configuramos el timeout de conexion a la BD
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//configura el nombre de la BD y la colleccion
	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	//se ejecutara la funcion DeleteOne, barraremos con el objeto Relacion un registro
	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		//retorna "false" si hubo un error y la descripcion
		return false, err
	}

	//retorna "true" si todo es OK y el error nulo
	return true, nil
}
