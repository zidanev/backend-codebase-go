package models

type User struct {
	Id           string `json:"id,omitempty" bson:"_id,omitempty"`
	Username     string `json:"username" bson:"username" validate:"required"`
	Password     string `json:"password" bson:"password" validate:"required"`
	Email        string `json:"email" bson:"email" validate:"required"`
	FullName     string `json:"fullName" bson:"fullName" validate:"required"`
	MobileNumber string `json:"mobileNumber" bson:"mobileNumber" validate:"required"`
	Status       string `json:"status,omitempty" bson:"status,omitempty"`
}

type UpsertUser struct {
	Username     string `json:"username,omitempty" bson:"username,omitempty"`
	Password     string `json:"password,omitempty" bson:"password,omitempty"`
	Email        string `json:"email,omitempty" bson:"email,omitempty"`
	FullName     string `json:"fullName,omitempty" bson:"fullName,omitempty"`
	MobileNumber string `json:"mobileNumber,omitempty" bson:"mobileNumber,omitempty"`
	Status       string `json:"status,omitempty" bson:"status,omitempty"`
}

func (u User) UpsertUser() UpsertUser {
	return UpsertUser{
		Username:     u.Username,
		Password:     u.Password,
		Email:        u.Email,
		FullName:     u.FullName,
		MobileNumber: u.MobileNumber,
		Status:       u.Status,
	}
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Id          string `json:"id"`
	AccessToken string `json:"accessToken"`
}

type GetUserResponse struct {
	Id           string `json:"id,omitempty"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	FullName     string `json:"fullName"`
	MobileNumber string `json:"mobileNumber"`
	Status       string `json:"status,omitempty"`
}
