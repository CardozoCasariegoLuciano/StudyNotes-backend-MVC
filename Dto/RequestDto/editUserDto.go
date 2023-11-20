package requestDto

type EditUserDto struct {
	Name  string `json:"name" validate:"min=5"`
	Image string `json:"image"`
}
