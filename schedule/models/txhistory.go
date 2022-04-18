package models

import (
	"pledge-bridge-backend/db"
	"pledge-bridge-backend/log"
	"pledge-bridge-backend/utils"
	"time"
)

type TxHistory struct {
	Id          int64     `json:"-" gorm:"column:id;primaryKey;AUTO_INCREMENT"`
	TxType      string    `json:"tx_type" gorm:"column:tx_type"` //deposit（0） or withdraw（1）
	SrcChain    string    `json:"src_chain" gorm:"column:src_chain"`
	DestChain   string    `json:"dest_chain" gorm:"column:dest_chain"`
	Asset       string    `json:"asset" gorm:"column:asset"`
	Amount      string    `json:"amount" gorm:"column:amount"`
	Fee         string    `json:"fee" gorm:"column:fee"`
	ActionAt    time.Time `json:"action_at" gorm:"column:action_at"`
	Address     string    `json:"address" gorm:"column:address"`
	DepositHash string    `json:"deposit_hash" gorm:"column:deposit_hash"`
	BridgeHash  string    `json:"bridge_hash" gorm:"column:bridge_hash"`
	Status      int       `json:"status" gorm:"column:status"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func NewTxHistory() *TxHistory {
	return &TxHistory{}
}

func (m *TxHistory) UpdateBridgeHash(whereCondition, hash, fee string) error {

	err := db.Mysql.Table("txhistory").Where(whereCondition).Updates(map[string]interface{}{
		"bridge_hash": hash,
		"updated_at":  utils.NowDataTime(),
	}).Debug().Error

	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}

	return nil
}
