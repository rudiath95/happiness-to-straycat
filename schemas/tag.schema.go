package schemas

type CreateTag struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTag struct {
	Name string `json:"name"`
}
