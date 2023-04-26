package bd

import (
	"context"
	"time"

	"github.com/xfred19x/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoUsuariosTodos Lee los usuarios registrados en el sistema, si se recibe "R", trae solo los que se relacionan conmigo*/
func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {

	//condigurar el tiemout de conexion a la BD
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//condigura el nombre de la BD y la colleccion
	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	//se crea un slice vacio del tipo usuario
	var results []*models.Usuario

	//configuro las propiedades que tendra mi consulta con option.Find
	//la paginacion y el limite de ellos
	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	//creamos la sentencia con la palabra a consultar en "search"
	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	//consulta en mongo BD con la funcion Find() con la sentencia y las propiedades de consulta
	cur, err := col.Find(ctx, query, findOptions)
	//valida si existe error
	if err != nil {
		//en caso de error retorna resultado en vacio y el status "false"
		return results, false
	}

	//se crea las variables que sera nuestros indicadores de encontrado o de incluirlo en el slice de respuesta
	var encontrado, incluir bool

	for cur.Next(ctx) {

		//creamos un objeto model usuario vacio
		var s models.Usuario

		//aqui asignamos un registro del cursor en el objeto
		err := cur.Decode(&s)
		//validamos si hay error en la decodificacion
		if err != nil {
			//en caso de error retorna resultado en vacio y el status "false"
			return results, false
		}

		//creamos un objeto models Relacion vacio y le asignamos los parametros de relacion
		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()

		//inicializamos la variable si lo va incluir o no
		incluir = false

		//llamamos la rutina consulto relacion
		encontrado, _ = ConsultoRelacion(r)

		//incluiremos si quiero listar usuarios nuevos que no tiene relacion conmigo
		if tipo == "new" && !encontrado {
			incluir = true
		}
		//incluiremos si solo queremos listar a los usuarios que sigo o tengan relacion conmigo
		if tipo == "follow" && encontrado {
			incluir = true
		}
		//no puedo incluir usuarios que tengan mi mismo ID
		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		//si llego incluir debo limpiar los campos sensibles de informacion
		if incluir {
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""

			//agregamos el modelo usuario al slice resultado
			results = append(results, &s)
		}
	}

	//validar si en el cursos hay algun error interno
	err = cur.Err()
	//validar si hubo error
	if err != nil {
		//en caso de error retorna resultado en vacio y el status "false"
		return results, false
	}

	//cerramos el cursor
	cur.Close(ctx)
	//retornar el slice cargado y true si todo es OK
	return results, true
}
