package response

type AuthResponse struct {
	JWT     string `json:"jwt,omitempty"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
