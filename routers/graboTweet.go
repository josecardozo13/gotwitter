package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/josecardozo13/gotwitter/bd"
	"github.com/josecardozo13/gotwitter/models"
)

/*GraboTweet permite grabar el tweet en la BD*/
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var t models.Tweet

	err := json.NewDecoder(r.Body).Decode(&t)

	registro := models.GraboTweet{
		UserId:  IDUsuario,
		Mensaje: t.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro. Reintente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el tweet", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
