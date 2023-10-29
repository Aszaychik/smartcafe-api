package web

type WebResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type AdminLoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}