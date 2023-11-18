package responseDto

type UserDto struct {
	ID    uint   `json:"id" mapper:"_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"image"`
	Role  string `json:"role"`
}
