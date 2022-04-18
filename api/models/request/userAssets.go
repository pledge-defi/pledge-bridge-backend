package request

type UserAssets struct {
	Token string `form:"token" json:"token" binding:"required"`
}
