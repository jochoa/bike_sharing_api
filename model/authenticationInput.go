package model

type AuthenticationInput struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
	Role uint `json:"role"`
}