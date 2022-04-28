package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/josecardozo13/gotwitter/bd"
)

/*ListaUsuarios devuelve la lista de usuarios*/
func ListaUsuarios(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el param pagina mayor a cero.", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := bd.LeoUsuarioTodos(IDUsuario, pag, search, typeUser)

	if status == false {
		http.Error(w, "Error al leer los usuarios.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
