package bd

import (
	"context"
	"log"

	//driver para conectarme a mongodb
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexión a la BD */
var MongoCN = ConectarBD()

// usaremos la funcion "Client" es del paquete mongo que sirve para agregar la cadena de conexion a la BD
var clientOptions = options.Client().ApplyURI("mongodb+srv://fgonzales:ZzXsgMxx6e4kH8u3@twitter.hgy5riu.mongodb.net/?retryWrites=true&w=majority")

/*ConectarBD es la función que me permite conectar la BD */
// retorna un objeto del tipo "mongo.Client"
func ConectarBD() *mongo.Client {

	//hace una conexion a la BD
	//el contexto es un espacio de memoria que puedo ir compartiendo a lo largo de la ejecucion
	//cuando usas el "context.TODO()" significa que se conecta a la BD sin ningun tipo de timeout u otra configuracion mas que default
	//tambien se puede crear variables y que esten disponibles en todo el contexto
	client, err := mongo.Connect(context.TODO(), clientOptions)

	//validamos si hay un error dutante la conexion
	if err != nil {
		//se mostrara por la consola de error el fatal
		//para mostrar el detalle del error en string se usa err.Error()
		log.Fatal(err.Error())
		//si hay error, se tiene que retornar el "client" asi este vacío
		return client
	}

	//esto es para ver si la base de datos esta arriba
	err = client.Ping(context.TODO(), nil)

	//validara si hubo algun error con la BD
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexión Exitosa con la BD")
	return client
}

/*ChequeoConnection es el Ping a la BD */
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
