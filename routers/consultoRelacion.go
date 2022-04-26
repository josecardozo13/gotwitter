package routers

import (
	"encoding/json"
	"net/http"

	"github.com/josecardozo13/gotwitter/bd"
	"github.com/josecardozo13/gotwitter/models"
)

/*ConsultoRelacion chequesa si hay una relacion entre dos usuarios*/
func ConsultoRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var response models.ResponseConsultaRelacion

	status, err := bd.ConsultoRelacion(t)
	if err != nil || status == false {
		response.Status = false
	} else {
		response.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
