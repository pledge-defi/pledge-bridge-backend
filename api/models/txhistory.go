package models

import (
	"pledge-bridge-backend/api/models/request"
	"pledge-bridge-backend/db"
	"pledge-bridge-backend/log"
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

func (m *TxHistory) Pagination(req *request.TxHistory, whereCondition string) (error, int64, []TxHistory) {
	var total int64
	var txHistorys []TxHistory

	db.Mysql.Table("txhistory").Where(whereCondition).Count(&total)

	err := db.Mysql.Table("txhistory").Where(whereCondition).Order("id desc").Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Find(&txHistorys).Debug().Error
	if err != nil {
		return err, 0, nil
	}

	return nil, total, txHistorys
}

func (m *TxHistory) AddRecord(req *request.AddTx) error {

	nowTime := time.Now()
	data := TxHistory{}

	data.TxType = req.TxType
	data.Amount = req.Amount
	data.Address = req.Address
	data.SrcChain = req.SrcChain
	data.DestChain = req.DestChain
	data.Asset = req.Asset
	data.Fee = req.Fee
	data.CreatedAt = nowTime
	data.UpdatedAt = nowTime
	data.ActionAt = nowTime
	data.DepositHash = req.TxHash

	err := db.Mysql.Table("txhistory").Create(&data)

	if err.Error != nil {
		log.Logger.Error(err.Error.Error())
		return err.Error
	}
	return nil
}
