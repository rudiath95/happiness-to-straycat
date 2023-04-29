package schemas

type CreateImmunization struct {
	Name string `json:"name" binding:"required"`
}

type UpdateImmunization struct {
	Name string `json:"name"`
}
