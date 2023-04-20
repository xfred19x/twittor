package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/xfred19x/twittor/middlew"
	"github.com/xfred19x/twittor/routers"
)

// Manejadores seteo mi puerto, el Handler y pongo a escuchar al Servidor
func Manejadores() {

	//Captura el HTTP para darle el manejo al response writer y al request, validara si hay informacion en el body o header
	router := mux.NewRouter()

	//Cuando desde el navegador se coloque "/registro" del tipo POST
	//este llamara a la funcion middleW con la funcion chequeoBD y si todo esta ok, retornara el control a routers
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	//obtiene el puerto si ya esta creado
	PORT := os.Getenv("PORT")
	if PORT == "" {
		//en caso no este creado, levantara el puerto 8080
		PORT = "8080"
	}

	//los cors son los permisos a mis Apis para que sean accesibles desde cualquier lugar
	//se usara la funcion AllowAll() para que le de permisto a cualquiera
	//Cuando al cors se le enviar un objeto router, a partir de ese momento el objeto cors toma el control de peticion de la web.
	handler := cors.AllowAll().Handler(router)
	//por si sucede algun error usaremos como traza el log.Fatal
	//con la funcion ListenAddServe, ponemos a escuchar el puerto y le pasemos el handlers creado.
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
