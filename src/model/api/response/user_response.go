package response

type UserLogin struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
