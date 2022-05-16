package models

/*LoginREsponse devuelve Token en el login*/
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
