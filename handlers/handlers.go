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

/*Manejadores seteo mi puerto, el Handler y pongo a escuchar al Servidor */
func Manejadores() {

	//Captura el HTTP para darle el manejo al response writer y al request, validara si hay informacion en el body o header
	router := mux.NewRouter()

	//Cuando desde el navegador se coloque "/registro" del tipo POST
	//este llamara la funcion middleW con la funcion chequeoBD y si todo esta ok, retornara el control a routers
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	//este llamara la funcion Login y si todo esta ok, retornara el control a routers
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")

	//este llamara la funcion verPerfil
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")

	//este llamara la funcion modificarPerfil
	router.HandleFunc("/modificarperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")

	//este llamara la funcion tweet
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")

	//este llamara la funcion leoTweets
	router.HandleFunc("/leotweets", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")

	//este llamara la funcion eliminarTweets
	router.HandleFunc("/eliminartweets", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")

	//este llamara la funcion subirAvatar
	router.HandleFunc("/subiravatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")

	//este llamara la funcion obtenerAvatar
	router.HandleFunc("/obteneravatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.ObtenerAvatar))).Methods("GET")

	//este llamara la funcion subirBanner
	router.HandleFunc("/subirbanner", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")

	//este llamara la funcion obtenerBanner
	router.HandleFunc("/obtenerbanner", middlew.ChequeoBD(middlew.ValidoJWT(routers.ObtenerBanner))).Methods("GET")

	//este llamara la funcion altaRelacion
	router.HandleFunc("/altarelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.AltaRelacion))).Methods("GET")

	//obtiene el puerto si ya esta creado
	PORT := os.Getenv("PORT")
	if PORT == "" {
		//en caso no este creado, levantara el puerto 8080
		PORT = "8081"
	}

	//los cors son los permisos a mis Apis para que sean accesibles desde cualquier lugar
	//se usara la funcion AllowAll() para que le de permisto a cualquiera
	//Cuando al cors se le enviar un objeto router, a partir de ese momento el objeto cors toma el control de peticion de la web.
	handler := cors.AllowAll().Handler(router)
	//por si sucede algun error usaremos como traza el log.Fatal
	//con la funcion ListenAddServe, ponemos a escuchar el puerto y le pasemos el handlers creado.
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
