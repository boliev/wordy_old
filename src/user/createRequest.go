package user

// CreateRequest request DTO for creating user
type CreateRequest struct {
	Name     string
	Email    string
	Password string
}
