package responseDto

type ResponseDto struct {
	Message     string      `json:"message"`
	MessageType string      `json:"messageType"`
	Data        interface{} `json:"data"`
}

func NewResponse(messageType string, message string, data interface{}) ResponseDto {
	return ResponseDto{message, messageType, data}
}
