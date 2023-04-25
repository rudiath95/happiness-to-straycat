package schemas

type CreateUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
