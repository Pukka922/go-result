package goresult

func Create[T any](success bool, data T, messages []Message) Result[T] {
	return Result[T]{
		Success:  success,
		Data:     data,
		Messages: messages,
	}
}

func Success[T any](data T) Result[T] {
	return Result[T]{
		Success: true,
		Data:    data,
	}
}

func SuccessWithMessage[T any](data T, message string) Result[T] {
	result := Result[T]{
		Success: true,
		Data:    data,
	}

	newMessage := CreateMessage(MessageTypeSuccess, message)
	result.Messages = append(result.Messages, newMessage)

	return result
}

func Error[T any](message string) Result[T] {
	result := Result[T]{
		Success: false,
	}

	newMessage := CreateMessage(MessageTypeError, message)
	result.Messages = append(result.Messages, newMessage)

	return result
}

func Info[T any](data T, message string) Result[T] {
	result := Result[T]{
		Success: true,
		Data:    data,
	}

	newMessage := CreateMessage(MessageTypeInfo, message)
	result.Messages = append(result.Messages, newMessage)

	return result
}

func Warning[T any](data T, message string) Result[T] {
	result := Result[T]{
		Success: true,
		Data:    data,
	}

	newMessage := CreateMessage(MessageTypeWarning, message)
	result.Messages = append(result.Messages, newMessage)

	return result
}
