package request

type TxHistory struct {
	//Address  string `form:"address" json:"address" binding:"omitempty"`
	Address  string `json:"address"`
	TxType   string `json:"txType" `
	Page     int    `json:"page" `
	PageSize int    `json:"pageSize" binding:"numeric"`
}
