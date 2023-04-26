package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/xfred19x/twittor/bd"
	"github.com/xfred19x/twittor/models"
)

/*SubirBanner sube el Banner al servidor */
func SubirBanner(w http.ResponseWriter, r *http.Request) {

	//vamos a capturar del request el archivo
	file, handler, err := r.FormFile("banner")

	//validamos si hubo algun error al obtener el archivo
	if err != nil {
		http.Error(w, "Error con el archivo "+err.Error(), 400)
		return
	}

	//del archivo obtendremos el nombre de la extension del archivo
	var extension = strings.Split(handler.Filename, ".")[1]

	//Definimos donde se va guardar el archivo, es muy importante que las carpetas existan, sino no lo guardará
	//tendra la estructura IDUsuario.extension del archivo
	var archivo string = "uploads/banners/" + IDUsuario + "." + extension

	//Aqui creamos el archivo en el disco, con los atributos y permisos: 0666 es el permiso de lectura y escritura
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	//si hubo algun error
	if err != nil {
		http.Error(w, "Error al subir la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	//vamos a copiar el archivo ingresado con los atributos y nombre del archivo recien creado
	_, err = io.Copy(f, file)
	//validar si hubo error
	if err != nil {
		http.Error(w, "Error al copiar la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	//se crea la variable usuario en vacio para modificar el campo Banner
	var usuario models.Usuario
	var status bool

	usuario.Banner = IDUsuario + "." + extension
	//se llama la rutina ModificoRegistro para actualizar el Banner
	status, err = bd.ModificoRegistro(usuario, IDUsuario)
	//valida si ocurre un error
	if err != nil || !status {
		http.Error(w, "Error al grabar el banner en la BD ! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
