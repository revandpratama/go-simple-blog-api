package dto

type LoginRequest struct {
	EmailOrUsername string `json:"email"`
	Password        string `json:"password"`
}

type LoginResponse struct {
	ID    int
	Name  string
	Token string `json:"token"`
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Username string `json:"username" validate:"required,alphanum"`
	// Username string  `json:"username" validate:"required,alphanum,startwith=@,gtefield=Name"`
	Email           string `json:"email" validate:"required,email:rfc"`
	Password        string `json:"password" validate:"required,min=8,max=36"`
	PasswordConfirm string `json:"confirm_password" validate:"eqfield=Password"`
}

type RegisterResponse struct {
}
