package bd

import (
	"context"
	"time"

	"github.com/xfred19x/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*LeoTweetsSeguidores lee los tweets de mis seguidores */
func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {

	//configuramos el timeout de conexion
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//configuramos el nombre de la bd y colleccion
	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	//creamos una variable con los datos de paginacion
	skip := (pagina - 1) * 20

	//crearemos un slice donde agregaremos todos los filtros de condiciones
	condiciones := make([]bson.M, 0)
	//debemos filtrarlo por el Id que vino de parametro
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	//condicion tipo bson que permite unir 2 tablas
	//en "from" la tabla que quiero relacionarme
	//en "localField" el campo que tendra la relacion
	//en "foreignField" el campo de "tweet" que tendra el id
	//en "as" el alias que vamos a llamarla a la tabla, puede ser cualquiera pero aqui le ponen "twwet"
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	//se usa al "unwind" para que todos los documentos vengan igual
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	//se prdenara por fecha
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.fecha": -1}})
	//se paginara
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	//creamos el cursos con la funcion Aggregate y nuestras condiciones
	cursor, _ := col.Aggregate(ctx, condiciones)

	//creamos la variable result con el slice de seguidores en vacio
	var result []models.DevuelvoTweetsSeguidores

	//ejecuta y arma todo el  documento con el formato que requiere en result
	err := cursor.All(ctx, &result)
	//validar si hubo error
	if err != nil {
		//en caso de error retorna resultado en vacio y el status "false"
		return result, false
	}

	//retornar el slice cargado y true si todo es OK
	return result, true

}
