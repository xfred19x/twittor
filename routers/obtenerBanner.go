package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/xfred19x/twittor/bd"
)

/*ObtenerBanner envia el Banner al HTTP */
func ObtenerBanner(w http.ResponseWriter, r *http.Request) {

	//Obtenemos de la URL el parametro "id"
	ID := r.URL.Query().Get("id")
	//validamos que el "id" tenga contenido
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	//luego buscamos el perfil de la BD
	perfil, err := bd.BuscoPerfil(ID)
	//valida si tuvo error
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	//Abre el archivo de la carpeta correspondiente del usuario
	OpenFile, err := os.Open("uploads/banners/" + perfil.Banner)
	//valida si hay error
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}

	//Envia el imagen a la variable "w" que es http.ResponseWriter en binario
	_, err = io.Copy(w, OpenFile)
	//valida si hay error
	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
	}
}
