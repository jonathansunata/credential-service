package entity

import "github.com/SawitProRecruitment/UserService/domain"

type User struct {
	Id          int32  `json:"id"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Salt        string `json:"salt"`
	CountLogin  *int32 `json:"count_login"`
}

func (u *User) Merge(request *domain.UpdateUserRequest) {
	if request.FullName != "" {
		u.FullName = request.FullName
	}

	if request.PhoneNumber != "" {
		u.PhoneNumber = request.PhoneNumber
	}
}
