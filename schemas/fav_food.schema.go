package schemas

type CreateFavFood struct {
	Company string `json:"company" binding:"required"`
	Variety string `json:"variety" binding:"required"`
	Protein int64  `json:"protein" binding:"required"`
	Fat     int64  `json:"fat" binding:"required"`
	Carbs   int64  `json:"carbs" binding:"required"`
	Phos    int64  `json:"phos" binding:"required"`
	Notes   string `json:"notes" binding:"required"`
}

type UpdateFavFood struct {
	Company string `json:"company"`
	Variety string `json:"variety"`
	Protein int64  `json:"protein"`
	Fat     int64  `json:"fat"`
	Carbs   int64  `json:"carbs"`
	Phos    int64  `json:"phos"`
	Notes   string `json:"notes"`
}
