package user

// Repository for user package
type Repository interface {
	Save(user User) User
}
