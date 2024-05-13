package goresult

func SimpleCreate(success bool, messages []Message) SimpleResult {
	return SimpleResult{
		Success:  success,
		Messages: messages,
	}
}

func SimpleSuccess() SimpleResult {
	return SimpleResult{
		Success: true,
	}
}

func SimpleSuccessWithMessage(message string) SimpleResult {
	result := SimpleResult{
		Success: true,
	}

	newMessage := CreateMessage(MessageTypeSuccess, message)
	result.Messages = append(result.Messages, newMessage)

	return result
}

func SimpleError(message string) SimpleResult {
	result := SimpleResult{
		Success: false,
	}

	newMessage := CreateMessage(MessageTypeError, message)
	result.Messages = append(result.Messages, newMessage)

	return result
}

func SimpleInfo(message string) SimpleResult {
	result := SimpleResult{
		Success: true,
	}

	newMessage := CreateMessage(MessageTypeInfo, message)
	result.Messages = append(result.Messages, newMessage)

	return result
}

func SimpleWarning(message string) SimpleResult {
	result := SimpleResult{
		Success: true,
	}

	newMessage := CreateMessage(MessageTypeWarning, message)
	result.Messages = append(result.Messages, newMessage)

	return result
}
