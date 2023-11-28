package domain

type UpdateUserRequest struct {
	PhoneNumber string `json:"phone_number"`
	FullName    string `json:"full_name"`
}
