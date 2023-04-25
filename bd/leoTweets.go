package bd

import (
	"context"
	"time"

	"github.com/xfred19x/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoTweets lee los tweets de un perfil */
func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {

	//configuramos el timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//configuramos el nombre de la BD y colleccion
	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	//creamos el slice con la estructura del model DevuelvoTweets en vacio
	var resultados []*models.DevuelvoTweets

	//creamos la condicion a buscar en la BD
	condicion := bson.M{
		"userid": ID,
	}

	//Usamos la funcion "Find" de options
	//aqui podemos configurar las propiedades de filtro segun la condicion
	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	//aqui establecemos la cantidad de registro que tendra por paginas, para el ejemplo sera 20
	opciones.SetSkip((pagina - 1) * 20)

	//aqui buscamos con la condicion y propiedades de option los registros grabados en el cursor
	cursor, err := col.Find(ctx, condicion, opciones)
	//validamo si ocurrio un error
	if err != nil {
		//retornamos resultados en vacio y false porque hubo error
		return resultados, false
	}

	//recorremos el cursor
	for cursor.Next(context.TODO()) {
		//creamos el objeto DevuelvoTweets para poder llenar el slice de resultados
		var registro models.DevuelvoTweets
		//del cursor grabo un registro en el objeto
		err := cursor.Decode(&registro)
		//validamos si hay error
		if err != nil {
			//retornamos resultados en vacio y false porque hubo error
			return resultados, false
		}
		//si todo es ok, se agrega el registro a resultados
		resultados = append(resultados, &registro)
	}

	//retornamos resultados de los mensajes tweets con todas las propiedades y filtros y true del proceso OK
	return resultados, true
}
