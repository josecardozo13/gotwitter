package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/josecardozo13/gotwitter/bd"
)

/*ObtenerBanner envia el banner al HTTP*/
func ObtenerBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el usuario "+err.Error(), http.StatusBadRequest)
		return
	}
	OpenFile, err := os.Open("uploads/banners/" + perfil.Banner)
	if err != nil {
		http.Error(w, "Banner no encontrado "+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar el banner "+err.Error(), http.StatusBadRequest)
	}
}
