package requestDto

type RegisterUserDto struct {
	Name         string `json:"name" validate:"required,min=3"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=3"`
	Confirmation string `json:"confirmation" validate:"required,eqfield=Password"`
}
