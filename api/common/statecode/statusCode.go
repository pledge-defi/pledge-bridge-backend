package statecode

const (
	// LangZh language
	LangZh   = 111
	LangEn   = 112
	LangZhTw = 113

	// CommonSuccess common
	CommonSuccess      = 0
	CommonErrServerErr = 1000
	ParameterEmptyErr  = 1001

	// TokenEmpty token
	TokenEmpty = 1101 //token empty
	TokenErr   = 1102 //token error

	// AddressEmpty txHistory
	AddressEmpty = 1201 //address empty

	// TxTypeEmpty addTx
	TxTypeEmpty    = 1301 //txType empty
	AssetEmpty     = 1302 //asset empty
	TxHashEmpty    = 1303 //txHash empty
	AmountEmpty    = 1304 //amount empty
	FeeEmpty       = 1305 //fee empty
	SrcChainEmpty  = 1306 //src_chain empty
	DestChainEmpty = 1307 //dest_chain empty

)

var Msg = map[int]map[int]string{

	0: {
		LangZh:   "成功",
		LangZhTw: "成功",
		LangEn:   "success",
	},
	1000: {
		LangZh:   "服务器繁忙，请稍后重试",
		LangZhTw: "服務器繁忙，請稍後重試",
		LangEn:   "server is busy, please try again later",
	},
	1001: {
		LangZh:   "参数不能为空",
		LangZhTw: "参数不能為空",
		LangEn:   "parameter is empty",
	},
	1101: {
		LangZh:   "token 不能为空",
		LangZhTw: "token 不能為空",
		LangEn:   "token required",
	},
	1102: {
		LangZh:   "token错误",
		LangZhTw: "token錯誤",
		LangEn:   "token invalid",
	},
	1201: {
		LangZh:   "address 不能为空",
		LangZhTw: "address 不能為空",
		LangEn:   "address required",
	},
	1301: {
		LangZh:   "txType 不能为空",
		LangZhTw: "txType 不能為空",
		LangEn:   "txType required",
	},
	1302: {
		LangZh:   "asset 不能为空",
		LangZhTw: "asset 不能為空",
		LangEn:   "asset required",
	},
	1303: {
		LangZh:   "txHash 不能为空",
		LangZhTw: "txHash 不能為空",
		LangEn:   "txHash required",
	},
	1304: {
		LangZh:   "amount 错误",
		LangZhTw: "amount 錯誤",
		LangEn:   "amount error",
	},
	1305: {
		LangZh:   "fee 错误",
		LangZhTw: "fee 錯誤",
		LangEn:   "fee error",
	},
	1306: {
		LangZh:   "src_chain 错误",
		LangZhTw: "src_chain 錯誤",
		LangEn:   "src_chain error",
	},
	1307: {
		LangZh:   "dest_chain 错误",
		LangZhTw: "dest_chain 錯誤",
		LangEn:   "dest_chain error",
	},
}

func GetMsg(c int, lang int) string {
	_, ok := Msg[c]
	if ok {
		msg, ok := Msg[c][lang]
		if ok {
			return msg
		}
		return Msg[CommonErrServerErr][lang]
	}
	return Msg[CommonErrServerErr][lang]
}
