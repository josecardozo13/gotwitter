package models

type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioid"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuariorelacionid"`
}

type ResponseConsultaRelacion struct {
	Status bool `json:"status"`
}
