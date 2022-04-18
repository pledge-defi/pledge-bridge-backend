package models

type UserAssets struct {
	Id               int    `json:"-" gorm:"column:id;primaryKey"`
	Token            string `json:"token" gorm:"column:token"`
	LockedPlgr       string `json:"locked_plgr" gorm:"column:locked_plgr"`
	PlgrCanWithdraw  string `json:"plgr_can_withdraw" gorm:"column:plgr_can_withdraw"`
	MplgrCanWithdraw string `json:"mplgr_can_withdraw" gorm:"column:mplgr_can_withdraw"`
	CreatedAt        string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt        string `json:"updated_at" gorm:"column:updated_at"`
}
