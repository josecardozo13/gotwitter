package routers

import (
	"encoding/json"
	"net/http"

	"github.com/josecardozo13/gotwitter/bd"
	"github.com/josecardozo13/gotwitter/models"
)

/*ModificarPerfil modifica el perfil de usuario*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}

	status, err := bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro. Reintente nuevamente"+err.Error(), 400)
	}

	if status == false {
		http.Error(w, "No se ha logrado modificar el registro del usuario", 400)
	}
	w.WriteHeader(http.StatusCreated)
}
