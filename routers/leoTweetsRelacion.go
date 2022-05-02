package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/josecardozo13/gotwitter/bd"
)

/*LeoTweetSeguidores lee los tweets de todos nuestros seguidores*/
func LeoTweetSeguidores(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, correcto := bd.LeoTweersSeguidores(IDUsuario, pagina)
	if correcto == false {
		http.Error(w, "Error al leer los twwets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
