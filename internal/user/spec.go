package user

type CreateUserSpec struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}