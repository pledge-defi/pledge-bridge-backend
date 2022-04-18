package response

type UserAssets struct {
	Token            string `json:"token"`
	LockedPlgr       string `json:"locked_plgr"`
	PlgrCanWithdraw  string `json:"plgr_can_withdraw"`
	MplgrCanWithdraw string `json:"mplgr_can_withdraw"`
}
