package models

type CreateAccount struct {
	Username string      `json:"username" binding:"required"`
	Password string      `json:"password" binding:"required"`
	Role     AccountRole `json:"role" binding:"required"`
	Email    string      `json:"email" binding:"required"`
	Age      int         `json:"age" binding:"required"`
}

type LoginAccount struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}