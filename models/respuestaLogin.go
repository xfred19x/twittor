package models

/*RespuestaLogin tiene el token que se devuelve con el login*/
type RespuesatLogin struct {
	Token string `json:"token,omitempty"`
}
