package models

import (
	"errors"
	"gorm.io/gorm"
	"pledge-bridge-backend/api/models/request"
	"pledge-bridge-backend/db"
	"pledge-bridge-backend/utils"
)

type UserAssets struct {
	Id               int    `json:"-" gorm:"column:id;primaryKey"`
	Token            string `json:"token" gorm:"column:token"`
	LockedPlgr       string `json:"locked_plgr" gorm:"column:locked_plgr"`
	PlgrCanWithdraw  string `json:"plgr_can_withdraw" gorm:"column:plgr_can_withdraw"`
	MplgrCanWithdraw string `json:"mplgr_can_withdraw" gorm:"column:mplgr_can_withdraw"`
	CreatedAt        string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt        string `json:"updated_at" gorm:"column:updated_at"`
}

func NewUserAssets() *UserAssets {
	return &UserAssets{}
}

func (m *UserAssets) GetInfoByToken(req *request.UserAssets) (error, UserAssets) {

	var userAssets UserAssets

	err := db.Mysql.Table("user_assets").Where("token=?", req.Token).First(&userAssets).Debug().Error
	if err != nil {
		return err, userAssets
	}

	return nil, userAssets
}

func (m *UserAssets) UpdateUserAssets(token string, data map[string]interface{}) error {

	err := db.Mysql.Table("user_assets").Where("token=?", token).First(&m).Debug().Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			data["token"] = token
			data["created_at"] = utils.NowDataTime()
			err = db.Mysql.Table("user_assets").Create(&data).Debug().Error
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	err = db.Mysql.Table("user_assets").Where("token=?", token).Updates(data).Debug().Error
	if err != nil {
		return err
	}

	return nil
}
