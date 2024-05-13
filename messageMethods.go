package goresult

const messageTypeSuccessString = "success"
const messageTypeWarningString = "warning"
const messageTypeErrorString = "error"
const messageTypeInfoString = "info"

const (
	MessageTypeSuccess MessageType = iota
	MessageTypeWarning
	MessageTypeError
	MessageTypeInfo
)

type MessageType int

func CreateMessage(messageType MessageType, message string) Message {
	messageTypeValue := ""

	switch messageType {
	case MessageTypeSuccess:
		messageTypeValue = messageTypeSuccessString
	case MessageTypeWarning:
		messageTypeValue = messageTypeWarningString
	case MessageTypeError:
		messageTypeValue = messageTypeErrorString
	case MessageTypeInfo:
		messageTypeValue = messageTypeInfoString
	}

	return Message{
		Type: messageTypeValue,
		Text: message,
	}
}
