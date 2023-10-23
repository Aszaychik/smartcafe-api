package web

type AdminLoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}