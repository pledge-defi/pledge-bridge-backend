package request

type AddTx struct {
	Address   string `form:"address" json:"address" binding:"required"`
	TxType    string `form:"txType" json:"txType"  binding:"oneof=0 1"`
	Asset     string `form:"asset" json:"asset"  binding:"required"`
	SrcChain  string `json:"src_chain"  binding:"required"`
	DestChain string `json:"dest_chain"  binding:"required"`
	Fee       string `form:"fee" json:"fee"  binding:"required"`
	TxHash    string `form:"txHash" json:"txHash"  binding:"required"`
	Amount    string `form:"amount" json:"amount"  binding:"required"`
}
