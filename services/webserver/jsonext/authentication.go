package jsonext

import (
	"github.com/ungame/go-signup/pb/auth"
	"time"
)

type CreateAuthenticationInput struct {
	Email    string `json:"email,omitempty"    validate:"required,email,gte=1,lte=50" `
	Username string `json:"username,omitempty" validate:"required,gte=1,lte=20" `
	Password string `json:"password,omitempty" validate:"required,gte=3,lte=100" `
	Phone    string `json:"phone,omitempty"    validate:"required,gte=1,lte=15" `
}

func (i *CreateAuthenticationInput) ToProto() *auth.CreateAuthenticationRequest {
	return &auth.CreateAuthenticationRequest{
		Email:    i.Email,
		Username: i.Username,
		Password: i.Password,
		Phone:    i.Phone,
	}
}

type AuthenticationUserOutput struct {
	Id        string    `json:"id,omitempty"`
	Email     string    `json:"email,omitempty"`
	Username  string    `json:"username,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func NewAuthenticationUserOutputFromProto(p *auth.AuthenticationUser) *AuthenticationUserOutput {
	return &AuthenticationUserOutput{
		Id:        p.Id,
		Email:     p.Email,
		Username:  p.Username,
		Phone:     p.Phone,
		CreatedAt: p.CreatedAt.AsTime(),
		UpdatedAt: p.UpdatedAt.AsTime(),
	}
}
