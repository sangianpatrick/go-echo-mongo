package wrapper

// ResponseError a helper to wrapping error data for response
var ResponseError = map[string]interface{}{
	"error":   true,
	"data":    nil,
	"message": "",
}

// ResponseSuccess a helper to wrapping success data for response
var ResponseSuccess = map[string]interface{}{
	"error":   false,
	"data":    "",
	"message": "",
}
