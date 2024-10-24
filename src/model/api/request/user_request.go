package request

type UserRegister struct {
	Username string `json:"username" binding:"required,min=4,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required,min=4,max=100"`
	Password string `json:"password" binding:"required,min=6"`
}
