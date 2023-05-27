package registerAuth

type InputRegister struct {
	FirstName string `json:"firstname" validate:"required,lowercase"`
	LastName  string `json:"lastname" validate:"required,lowercase"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,gte=8"`
}
