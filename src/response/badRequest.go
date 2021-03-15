package response

// BadRequest struct
type BadRequest struct {
	Error string
}

// CreateBadRequest creates Bad Request body
func CreateBadRequest(errorMessage string) BadRequest {
	return BadRequest{Error: errorMessage}
}
