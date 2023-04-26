package bd

import (
	"context"
	"time"

	"github.com/xfred19x/twittor/models"
)

/*InsertoRelacion graba la relación en la BD */
func InsertoRelacion(t models.Relacion) (bool, error) {

	//configura timeout de conexion a la BD
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//configura el nombre de la BD y la colleccion
	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	//Inserta un registro con el objeto relacion en la BD
	_, err := col.InsertOne(ctx, t)
	//valida si hubo error
	if err != nil {
		//retorna en caso de error retorna "false" y le descripcion del error
		return false, err
	}

	//retorna "true" si todo el proceso fue OK y el error en nulo
	return true, nil
}
