package responseDto

type UserDto struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"image"`
	Role  string `json:"role"`
}
