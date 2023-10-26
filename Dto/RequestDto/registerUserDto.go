package requestDto

type RegisterUserDto struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Confirmation string `json:"confirmation"`
}
