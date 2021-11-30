package contracts
const (
	defaultMessage = ""
	defaultTitle = ""
)
type MessageContract struct {
	Type     MessageType     `json:"type"`
	Position MessagePosition `json:"position"`
	Content  string          `json:"content"`
	Title    string          `json:"title"`
	Code     string          `json:"code"`
}

type MessageType string
const(
	Information MessageType = "information"
	Warning MessageType = "warning"
	Error MessageType = "error"
)

type MessagePosition string
const (
	Top MessagePosition = "top"
	Bottom MessagePosition = "bottom"
	Left MessagePosition = "left"
	Right MessagePosition = "right"
)

func NewBadRequestErrorMessage(code string, message string) *MessageContract{
	return &MessageContract{
		Type:    Error,
		Position: Top,
		Content:  message,
		Title:    defaultTitle,
		Code:     code,
	}
}


func NewInternalServerErrorMessage(code string) *MessageContract{
	return &MessageContract{
		Type:    Error,
		Position: Top,
		Content:  defaultMessage,
		Title:    defaultTitle,
		Code:     code,
	}
}

